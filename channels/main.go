package main

import (
	"fmt"
	"net/http"
	"time"
)

// call every url in series
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// create a channel that works with strings
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// We can wait for the channels in a for loop.
	// "range c" means wait for the channel until
	// it receives a message then continue with the
	// body of the for loop.
	for l := range c {
		// This is a function literal. It is the equivalent
		// of a javascript IIFE
		// Also, if we don't explicitly pass the link to the
		// function literal, go will pass the reference to that
		// var instead of the actual value and the reference is changed
		// over time by the for loop
		go func(link string) {
			time.Sleep(time.Second)
			checkLink(link, c)
			// here, when we pass (l) to the function literal go will
			// make a copy of the value and pass the value instead of
			// passing the reference which is what would happen if we
			// didn't explicitly pass the link into the function literal
			// The lesson here is to never share the same variable
			// between 2 go routines. Always pass it to a function first
			// so a copy is created or pass it through a channel
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}

// when we prefix a func with the keyword "go", it'll spawn
// a new go routine and pass it to the scheduler and
// continue the execution of the main go routine (the main)
// program. The scheduler is then in charge of switching
// between go routines. When the main program finishes
// it exists even if there are still go routines going on.

// Channels are used to communicate between go routines.
// Channels are typed. We can only send a specific type
// of data over one channel.
