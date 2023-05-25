package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Make() will create a byte slcie of size 99999.
	// A slice can grow and shrink but we can also
	// initialise it with n elements
	// bs := make([]byte, 99999)

	// we pass the empty slice to Read. Then, Read will
	// mutate that slice and put the response data in it.
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// faster alternative to all the code above
	io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	// the Writer function needs to return the number of processed bytes as the first argument
	// and the error or nil as the second argument
	return len(bs), nil
}
