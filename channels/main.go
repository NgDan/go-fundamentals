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

	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		return
	}

	fmt.Println(link, "is up")
}

// when we prefix a func with the keyword "go", it'll spawn
// a new go routine and pass it to the scheduler and
// continue the execution of the main go routine (the main)
// program. The scheduler is then in charge of switching
// between go routines. When the main program finishes
// it exists even if there are still go routines going on.

// Channels are used to communicate between go routines.
