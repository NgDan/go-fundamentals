package main

import "fmt"

func main() {
	// the first string type is the type of the key
	// the second string type is the type of the value
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	// Maps properties can't be accessed using the dot syntax like structs.
	// They can only be accessed using square brackets syntax:
	// colors["red"] works but colors.red doesn't

	fmt.Println(colors)
}
