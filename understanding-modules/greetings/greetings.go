package greetings

import "fmt"

// Hello returns a greeting for the named person.
/*
	GO_TIP: A function whose name starts with a capital letter can be called by a function not in the same package. These are called exported names and are accessible outside the package
*/
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
