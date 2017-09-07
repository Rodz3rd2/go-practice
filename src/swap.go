package main

import "fmt"

func main() {
	var a, b string = swap("Hello", "World")
	fmt.Println(a, b)
}

func swap(a, b string) (string, string) {
	return b, a
}