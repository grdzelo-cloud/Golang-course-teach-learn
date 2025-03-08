package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	log.SetPrefix("greetings:")

	// Request a greeting message.
	message, err := greetings.Hello(" ")
	// If error war returned, print it to console and
	// exit the program

	if err != nil {
		log.Fatal(err)
	}

	// If no error war returned, print the returned message
	// to the console.
	fmt.Println(message)
}
