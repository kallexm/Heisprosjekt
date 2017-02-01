package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
	//"io"
)

func CheckError(err error) {
	//Chech if it is a timeout error. Will hapen if the server does not responds  
	if nerr, ok := err.(net.Error); ok &&  nerr.Timeout(){
		fmt.Println("Error timeout: ", err)
	} else if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

func main() {
	//Sets up a TCP connector from net.Conn
	Conn, err := net.Dial("tcp","129.241.187.43:33546")
	CheckError(err)

	defer Conn.Close()
	i := 0
	bufRecive := make([]byte, 1024)
	for{
		//Chech to se if it gets any replay
		n, err := Conn.Read(bufRecive)
		CheckError(err)
		fmt.Println("Mesage length is: ", n)
		fmt.Println("Mesage is: ", string(bufRecive[0:n]))
		time.Sleep(time.Second * 2)
		//Sets the timeout to one seconds afther curent time
		err = Conn.SetDeadline(time.Now().Add(time.Second*1))
		CheckError(err)
		//Puts the msg in a buffer
		msg := strconv.Itoa(i)+"\x00"
		bufSend := []byte(msg)
		i++
		//Sends the mesage
		_, err = Conn.Write(bufSend)
		CheckError(err)
		fmt.Println("Test")


	}
	
}