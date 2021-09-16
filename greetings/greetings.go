package greetings

import (
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string, color string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. [%v]Welcome!", name, color)

	return message
}
