package main

import (
	. "fmt"
	"net"


)

func main(){
	b := []byte{}
	Ip := []byte{} 
	resiveAddr := net.UDPAddr{Ip,30000,""}

	connection1, _ := net.ListenUDP("udp6", &resiveAddr)
	
	len, Addr, _ := connection1.ReadFromUDP(b)
	Println("Message from: ")
	Println(Addr.IP)
	Println("Message is %i long", len)
	Println("\nMessage is:0 ")
	Println(b)
}