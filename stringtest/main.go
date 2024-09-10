package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string
	s4 := "test"
	s1 = "here" // Assignment after declaration

	var s2 string = "im o k" // Declaration and initialization
	str := "a,b,c,d"
	parts := strings.Split(str, ",")
	var s3 = "Go Lang" // Type inferred by Go

	const name = "Strings" // Constant string

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(name)
	fmt.Println(str)
	fmt.Println(parts)
}
