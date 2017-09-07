package main

import "fmt"

type Student struct {
	fname, lname, gender string
	age int
}

func main() {
	rodz := Student{"Rodrigo III", "Galura", "male", 23}
	ross := Student{"Lynrd", "Alquiroz", "male", 26}

	fmt.Println("Rodz fullname:", rodz.fname, rodz.lname)
	fmt.Println("Ross fullname:", ross.fname, ross.lname)
}