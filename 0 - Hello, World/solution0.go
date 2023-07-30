package main

import (
	"fmt"
)

func main() {
	// Declare a variable named 'input_string' to hold our input.
	var input_string string

	// Read a full line of input from stdin and save it to our variable, input_string.
	fmt.Scanln(&input_string)

	// Print a string literal saying "Hello, World." to stdout using fmt.Println.
	fmt.Println("Hello, World.")

	// Print the contents of input_string to stdout.
	fmt.Println(input_string)
}
