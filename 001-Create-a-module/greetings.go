package greetings

import "fmt"

func Hello(name string) string {
	var message string
	message = fmt.Sprintf("H1, $v. Welcome!", name)
	return message
}
