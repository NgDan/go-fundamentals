package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// properties can be assinged in the same order they are declared
	// in the struct definition
	// alex := person{"Alex", "Anderson"}

	// we can also pass them by property name, which is much safer
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}
