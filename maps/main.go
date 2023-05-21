package main

import "fmt"

func main() {
	// the first string type is the type of the key
	// the second string type is the type of the value
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	fmt.Println(colors)
}
