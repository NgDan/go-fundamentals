package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	fmt.Println("hello")
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

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
