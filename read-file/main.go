package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	providedFilePath := os.Args[1]

	file, e := os.Open(providedFilePath)

	if e != nil {
		fmt.Println(e.Error())
	}

	io.Copy(os.Stdout, file)
}
