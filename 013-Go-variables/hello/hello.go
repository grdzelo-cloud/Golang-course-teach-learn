package main

import (
	"fmt"
)

// we can init variables outside the function
var z int
var v int = 2
var n = 3

// j := 7

func main() {
	var student1 string = "John" // type is string
	var student2 = "jane"        // type is inferred
	x := 2                       // type is inferred
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println(`davit`)

	fmt.Println("-------------------------------------------")
	fmt.Println("Variable Declaration Without Initial Value")
	fmt.Println("-------------------------------------------")

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("-------------------------------------------")
	fmt.Println("Value Assignment After Declaration")
	fmt.Println("-------------------------------------------")

	var student3 string
	student3 = "John"
	fmt.Println(student3)

	fmt.Println("-------------------------------------------")
	fmt.Println("Difference Between var and :=")
	fmt.Println("-------------------------------------------")

	z = 1
	fmt.Println(z)
	fmt.Println(v)
	fmt.Println(n)

	// fmt.Println(j) error output
}
