package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i int

func rutine1() {
	for j := 0; j < 1000000; j++ {
		i += 1
	}
}

func rutine2() {
	for j := 0; j < 1000000; j++ {
		i -= 1
	}
}

func main() {
	i = 0
	runtime.GOMAXPROCS(runtime.NumCPU())

	go rutine1()
	go rutine2()

	time.Sleep(500 * time.Microsecond)

	Println("The value is: ", i)
}
