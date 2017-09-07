package main

import (
	"fmt"
	"bytes"
)

func main() {
	var b bytes.Buffer

	b.WriteString("Hello World")
	fmt.Println(b.String())
}