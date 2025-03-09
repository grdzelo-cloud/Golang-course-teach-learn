package main

import (
	"fmt"
)

func main() {
	stringFormattingVerbs()
}

func formattingVerbsForPrintf() {
	var i = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)    // 15.5
	fmt.Printf("%#v\n", i)   // 15.5
	fmt.Printf("%#v%%\n", i) // 15.5%
	fmt.Printf("%T\n", i)    // float

	fmt.Printf("%v\n", txt)  // Hello World!
	fmt.Printf("%#v\n", txt) // "Hello World!"
	fmt.Printf("%T\n", txt)  // string
}

func integerFormattingVerbs() {
	var i = 15

	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%+d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%O\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%#x\n", i)
	fmt.Printf("%4d\n", i)
	fmt.Printf("%-4d\n", i)
	fmt.Printf("%04d\n", i)
}

func stringFormattingVerbs() {
	var txt = "Hello"

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)
}
func booleanFormattingVerbs() {
	var i = true
	var j = false

	fmt.Printf("%t\n", i)
	fmt.Printf("%t\n", j)
}
func floatFormattingVerbs() {
	var txt = "Hello"

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)
}
