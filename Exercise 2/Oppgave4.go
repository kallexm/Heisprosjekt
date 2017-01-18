package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i int
var mutex chan int

func rutine1() {
	for j := 0; j < 1000000; j++ {
		mutex<- 1;
		i += 1
		<- mutex;
		
	}
}

func rutine2() {
	for j := 0; j < 999999; j++ {
		mutex<- 1;
		i -= 1
		<- mutex;
	}
}

func main() {
	i = 0
	runtime.GOMAXPROCS(runtime.NumCPU())
	mutex = make(chan int);
	go rutine1()
	<- mutex
	go rutine2()
	mutex<- 1
	

	time.Sleep(10000 * time.Microsecond)

	
	Println("The value is: ", i)
}
