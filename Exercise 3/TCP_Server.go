package main 

import (
	"fmt"
	"net"
	"os"
	"io"
	"time"
	//"strconv"
	//"bytes"
)

func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

 func main(){
 	// creates a Listener
 	ln, err := net.Listen("tcp", ":50002")
 	CheckError(err)
 	defer ln.Close()
 	Conn, err := net.Dial("tcp","129.241.187.43:33546")
	CheckError(err)

	defer Conn.Close()
	//i := 0
	bufRecive := make([]byte, 1024)
	n, err := Conn.Read(bufRecive)
	CheckError(err)
	fmt.Println("Mesage is: ", string(bufRecive[0:n]))
	msg := "Connect to: 129.241.187.149:50002\x00"
	bufSend := []byte(msg) 
	_, err = Conn.Write(bufSend)
 	for {
 		//Accepts a TCP conection, conn is a net.Conn
 		conn, err := ln.Accept()
 		CheckError(err)

 		fmt.Printf("Resived conection %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

 		go handelConnection(conn)
  	}

 }


 func handelConnection(conn net.Conn){
 	newMsg := "Hello World!\x00"
 	bufSend := []byte(newMsg)
 	_, err := conn.Write(bufSend)
 	CheckError(err)
 	for{
	 	buf := make([]byte,1024)
	 	//Reads mesage
	 	n, err := conn.Read(buf)
	 	fmt.Println("Test")
	 	//Chech if conection is lost
	 	if err == io.EOF {
	 		return
	 	}
	 	CheckError(err)
	 	fmt.Println("The mesage is: ", string(buf[0:n]), "from: ", conn.RemoteAddr())
	 	//Writes back
	 	_, err = conn.Write(buf[0:n])
	 	time.Sleep(time.Second * 5)	
	 }
 }

