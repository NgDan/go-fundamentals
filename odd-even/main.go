package main

import (
	"fmt"
	"strconv"
)

func main() {
	sliceOfInts := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range sliceOfInts {
		isNumberOdd := i%2 == 1
		if isNumberOdd {
			fmt.Println(strconv.Itoa(i) + " is " + "odd")
		} else {
			fmt.Println(strconv.Itoa(i) + " is " + "even")
		}
	}
}
