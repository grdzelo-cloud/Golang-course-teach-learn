package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Davit"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print teh returned map of
	// messages th the console

	fmt.Println(messages)
}
