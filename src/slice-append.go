package main

import "fmt"

func main() {
	var array []int

	fmt.Println(array)
	array = append(array, 1)
	fmt.Println(array)
	array = append(array, 2, 3)
	fmt.Println(array)
	array = append(array, 4, 5, 6)
	fmt.Println(array)
	array = append(array, 7, 8, 9, 10)
	fmt.Println(array)
}