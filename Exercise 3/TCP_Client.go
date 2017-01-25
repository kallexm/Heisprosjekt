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
	Conn, err := net.Dial("tcp","10.22.72.203:50002")
	CheckError(err)

	defer Conn.Close()
	i := 0
	for{
		//Sets the timeout to five seconds afther curent time
		err = Conn.SetDeadline(time.Now().Add(time.Second*1))
		CheckError(err)
		//Puts the msg in a buffer
		msg := strconv.Itoa(i)
		buf := []byte(msg)
		i++
		//Sends the mesage
		_, err = Conn.Write(buf)
		CheckError(err)
		//Chech to se if it gets any replay
		n, err := Conn.Read(buf)
		CheckError(err)
		fmt.Println("Mesage is: ", string(buf[0:n]))
		time.Sleep(time.Second * 2)

	}
	
}