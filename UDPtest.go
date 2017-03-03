package main

import (
	"fmt"
	"net"
	"time"
)

func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}

func main() {
	fmt.Println("Starting Connection procedure...")

	toAddress, err := net.ResolveUDPAddr("udp", "255.255.255.255:60001")
	CheckError(err)

	fromAddress, err := net.ResolveUDPAddr("udp", net.JoinHostPort(GetLocalIP(), "60002"))
	CheckError(err)

	connection, err := net.DialUDP("udp", fromAddress, toAddress)
	CheckError(err)


	defer connection.Close()
	receiveMsg := make([]byte, 1024)
	for {
		bCastMsg := []byte("101")
		_, err = connection.Write(bCastMsg)
		CheckError(err)

		connection.SetReadDeadline(time.Now().Add(4*time.Second))
		m, fromAddr, err := connection.ReadFromUDP(receiveMsg)
		fmt.Println("Timeout")
		time.Sleep(1*time.Second)
		if err != nil {
			fmt.Println("Timeout: ", err)
			break
		}else{
			fmt.Println(string(receiveMsg[0:m]), "from", fromAddr)
			time.Sleep(1*time.Second)
		}
	}
	connection.Close()


	fmt.Println("Starting listen procedure...")

	listenAddr, err := net.ResolveUDPAddr("udp", ":60001")
	CheckError(err)
	listenConn, err := net.ListenUDP("udp", listenAddr)
	CheckError(err)
	defer listenConn.Close()
	
	buf := make([]byte, 1024)
	for {
		//connection.SetReadDeadline(time.Now())
		n, remoteAddr, err := listenConn.ReadFromUDP(buf)
		CheckError(err)
		if msg := string(buf[0:n]); msg == "101" {
			fmt.Println("if", string(msg[0:n]))
			fromAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:60002")
			CheckError(err)
			fmt.Println(remoteAddr)
			answerConn, err := net.DialUDP("udp", fromAddr, remoteAddr)
			CheckError(err)
			_, err = answerConn.Write([]byte("Hello!\n"))
			CheckError(err)
			fmt.Println("done with if")
			answerConn.Close()
		}else{
			fmt.Println("else", string(msg[0:n]))
		}
	}
}


func GetLocalIP() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, address := range addrs {
        // check the address type and if it is not a loopback the display it
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}