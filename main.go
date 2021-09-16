package main

import (
	"github.com/jeyabalajis/go-greetings/greetings"
	"github.com/mitchellh/colorstring"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys", "green")
	colorstring.Println(message)
}
