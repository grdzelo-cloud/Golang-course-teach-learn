package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty string")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages

	messages := make(map[string]string)

	// Loop thought the received slice of names, calling
	// the Hello function to get a message for each name.

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// If the map. associate the retrieved message with
		// the name
		messages[name] = message
	}
	return messages, nil

}

func randomFormat() string {
	formats := []string{
		"Zdarova %v. Dzma!",
		"Shen %v, ar xar?!",
		"Ra ubneli xar %v",
	}
	return formats[rand.Intn(len(formats))]
}
