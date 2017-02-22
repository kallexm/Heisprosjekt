package main 

import(
	"fmt"
	//"io/ioutil"
	"os/exec"
	"net"
	//"strconv"
	"time"
	"encoding/binary"
)

/*
Det meste av kodne fungerer. Den eneste feilen jeg klare å
finne er at konverteringen fra en byte[] stream til int ikke fungerer
Programet starte med å lytte etter heartbeats fra andre prosesser
Hvis den finner det skall den fortsette å lyte og ta vare på daten fra
den siste heartbeten. Om den ikke får en heartbeat skall den starte en 
backup og begyne å sende ut sin egen heartbeat. 
*/

func count(n int) int{
	n = n+1
	fmt.Println(n)
	return n
}

func creatBacup() error{
	cmd := exec.Command("bash","-c","gnome-terminal -e \"go run main.go\"")
	_, err := cmd.Output()
	if err != nil{
		return err
	}
	return nil
}

func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}


func setUpUDPServer()*net.UDPConn{
	/* Lets prepare a address at any address at port 20003*/   
    listAddr,err := net.ResolveUDPAddr("udp",":20003")
    CheckError(err)


    ServerConn, err := net.ListenUDP("udp", listAddr)
    CheckError(err)
    return ServerConn
}

func setUpUDPClient()*net.UDPConn{
	//Sets up the address for the server
    ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:20003")
    CheckError(err)
 
    //Sets up the local address 
    LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:20001")
    CheckError(err)
 

    //Sets up a UDP conection to the server
    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    CheckError(err)
    return Conn
}

func main() {
	var i int64;
	i = 20
	//Setter opp server og ser om det er noen som sender en heart beat
	conn := setUpUDPServer()
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(4*time.Second))
	buf := make([]byte, 2048)
	fmt.Println("Venter på Heartbeat")
	for {
		n,/*add*/_,err := conn.ReadFromUDP(buf)
		conn.SetReadDeadline(time.Now().Add(4*time.Second))
		//Skjekker som ReadFromUDP har en error
		if err != nil {
			//Hvis erroern er noe annet en timout hånter den
			if e, ok := err.(net.Error); !ok || !e.Timeout(){
				CheckError(e)
			}
			//Hvis erroren er en timeout forlat inf for løkke
			break
		}
		i,n = binary.Varint(buf[0:n])
		//fmt.Println("i is: ",i)
	}
	conn.Close()
	//stater en backup 
	err := creatBacup()
	if err != nil{
		fmt.Println("I paniced")
		panic(err)
	}
	time.Sleep(time.Second*2)
	//Setter opp et UDP broadcast for å send sin egen heartbeat
	//heartbeaten er tallet den skall skriv ut, det blir også skrevet ut til skjer
	conn = setUpUDPClient()
	defer conn.Close()
	for {
		fmt.Println("Nr is: ", i)
		//creates masage 
	    buf := make([]byte, 1024)
	    _ = binary.PutVarint(buf,i)
	    //fmt.Println("Nr bytes: ", n)
	    i++
	    //sends msg to server
	    _,err = conn.Write(buf)
	    if err != nil {
	        fmt.Println(i, err)
	    }
	    time.Sleep(time.Second * 2)
	}
}