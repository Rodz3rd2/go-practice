package main

import "fmt"

type Student struct {
	fname, lname string
}

func main() {
	students := make(map[string]Student)
	students["rodz"] = Student{"Rodrigo III", "Galura"}
	students["ross"] = Student{"Lynrd", "Alquiroz"}

	fmt.Println(students)
}