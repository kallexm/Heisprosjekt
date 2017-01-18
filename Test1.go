package main

import ( 
	."fmt"
)


type rect struct {
    width int
    height int
}

func (r *rect) area() int {
    return r.width * r.height
}

func main() {
	i := 1
	Println("The value is: ", i)
	//var r rect;
}