package main

import "fmt"

func main() {
	fmt.Println(plus(20, 10))
	fmt.Println(minus(20, 10))
	fmt.Println(times(20, 10))
	fmt.Println(divide(20, 10))
}

func plus(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func times(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}