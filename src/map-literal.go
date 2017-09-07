package main

import "fmt"

type Student struct {
	fname, lname string
}

func main() {
	var students = map[string]Student{
		"rodz": {"Rodrigo III", "Galura"},
		"ross": {"Lynrd", "Alquiroz"},
	}

	fmt.Println(students)
}