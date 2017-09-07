package main

import "fmt"

func main() {
	var sum, n int = 0, 10

	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Println(sum)
}