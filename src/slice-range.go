package main

import "fmt"

func main() {
	primes := [5]int{2, 3, 5, 7, 11}

	for _, v := range primes {
		fmt.Printf("%d\n", v)
	}
}