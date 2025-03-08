package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greetings for the named person.
func Hello(name string) (string, error) {
	// if no name was given, return error with a message
	if name == "" {
		return name, errors.New("empty string")
	}
	// create a message using a random format

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Create randomFormat function witch is return one of greeting messages
// The returned message is selected at random
func randomFormat() string {
	// A slice of message formats/
	formats := []string{
		"Hi, %v. Welcome",
		"Great to see you, %v.",
		"Hail, %v! well met!",
	}
	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
