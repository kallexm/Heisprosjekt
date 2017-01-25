package main
 
import (
    "fmt"
    "net"
    "time"
    "strconv"
)
 
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}
 
func main() {
    //Sets up the address for the server
    ServerAddr,err := net.ResolveUDPAddr("udp","10.22.72.203:50002")
    CheckError(err)
 
    //Sets up the local address 
    LocalAddr, err := net.ResolveUDPAddr("udp", "10.22.72.203:50000")
    CheckError(err)
 

    //Sets up a UDP conection to the server
    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    CheckError(err)
    
    defer Conn.Close()

    i := 0
    for {
        //creates masage 
        msg := strconv.Itoa(i)
        i++
        buf := []byte(msg)
        //sends msg to server
        _,err = Conn.Write(buf)
        if err != nil {
            fmt.Println(msg, err)
        }
        //resives msg from server 
        buf = make([]byte, 1024)
        n,addr,err := Conn.ReadFrom(buf)
        fmt.Println("Received ",string(buf[0:n]), " from ",addr)

        if err != nil {
            fmt.Println("Error: ",err)
        } 
        time.Sleep(time.Second * 2)
    }
}