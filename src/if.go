package main

import (
	"fmt"
)

func main() {
	year := 2017

	if year % 2 == 0 {
		fmt.Println(year, "is leap year.")
	} else {
		fmt.Println(year, "is not leap year.")
	}
}