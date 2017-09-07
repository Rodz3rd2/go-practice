package main

import "fmt"
import "time"


func main () {
	today := time.Now().Weekday()

	fmt.Println(today + 2)
}