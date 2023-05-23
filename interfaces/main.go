package main

import "fmt"

type bot interface {
	// Any type in this program that implements a function called getGreeting
	// which returns a type string will also be considered of type "bot".
	// This is why we say that interfaces in go are implicit. Go will automatically
	// cast this interface to any type that has the same functions as the ones
	// in the interface
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	fmt.Println("hello")
	eb := englishBot{}
	sb := spanishBot{}

	// because the englishBot and spanishBot both implement a function named "getGreeting"
	// that returns type "bot", they're both also considered a type "bot". That's why
	// we can pass them to printGreeting
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// we can skip naming the argument in the receiver if we're not going to use it in the function. We can just leave the type (englishBot)
func (englishBot) getGreeting() string {
	// this has a DIFFERENT implementation
	// than spanishBot getGreeting
	return "hello"
}

// we can skip naming the argument in the receiver if we're not going to use it in the function. We can just leave the type (spanishBot)
func (spanishBot) getGreeting() string {
	// this has a DIFFERENT implementation
	// than englishBot getGreeting
	return "hola"
}
