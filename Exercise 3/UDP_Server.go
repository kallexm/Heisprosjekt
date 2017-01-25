package main
 
import (
    "fmt"
    "net"
    "os"
)
 
/* A Simple function to verify error */
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}
 
func main() {
    /* Lets prepare a address at any address at port 50002*/   
    listAddr,err := net.ResolveUDPAddr("udp",":50002")
    CheckError(err)


    ServerConn, err := net.ListenUDP("udp", listAddr)
    CheckError(err)
    defer ServerConn.Close()
 
    buf := make([]byte, 1024)
    for {
    	//reads resived msg
        n,addr,err := ServerConn.ReadFromUDP(buf)
        fmt.Println("Received ",string(buf[0:n]), " from ",addr)

        if err != nil {
            fmt.Println("Error: ",err)
        }

        //sends resived msg back. 
        _,err = ServerConn.WriteTo(buf[0:n],addr)
        if(err != nil){
        	fmt.Println(string(buf[0:n]),err)
        } 
    }
}