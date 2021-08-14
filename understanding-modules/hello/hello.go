/*
	GO TIP: code executed as an application must be in a main package
*/
package main

import (
	"fmt"
	"shivanikanal/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Shivani")
	fmt.Println(message)
}
