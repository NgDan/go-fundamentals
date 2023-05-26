package main

import (
	"fmt"
	"net/http"
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

	// we can wait for the channels in a for loop
	// this syntax means the for loop will run forever
	for {
		go checkLink(<-c, c)
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
