package main

import "fmt"
import "runtime"

func main() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Printf("OS X.\n")
		case "linux":
			fmt.Printf("Linux\n")
		default:
			fmt.Printf("%s.\n", os)
	}
}