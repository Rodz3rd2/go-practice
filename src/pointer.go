package main

import "fmt"

func main() {

	x := 20

	p := &x // point x
	fmt.Println(*p) // print x through the pointer p
	*p = 21 // change the value of x through pointer p
	fmt.Println(x)
	fmt.Println(*p)
}