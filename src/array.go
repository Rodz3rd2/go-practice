package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
	primes = primes[2:5]
	fmt.Printf("len=%d cap=%d %v\n", len(primes), cap(primes), primes)
}