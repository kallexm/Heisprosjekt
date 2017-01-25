package main 

import (
	"fmt"
	"net"
	"os"
	"io"
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
 	for {
 		//Accepts a TCP conection, conn is a net.Conn
 		conn, err := ln.Accept()
 		CheckError(err)

 		fmt.Printf("Resived conection %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

 		go handelConnection(conn)
  	}

 }


 func handelConnection(conn net.Conn){
 	for{
	 	buf := make([]byte,1024)
	 	//Reads mesage
	 	n, err := conn.Read(buf)
	 	//Chech if conection is lost
	 	if err == io.EOF {
	 		return
	 	}
	 	CheckError(err)
	 	fmt.Println("The mesage is: ", string(buf[0:n]), "from: ", conn.RemoteAddr())
	 	//Writes back
	 	_, err = conn.Write(buf[0:n])	
	 }
 }

