package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jimparty@gmail.com",
			zipCode: 94000,
		},
	}
	// jim.updateName("Jimmy")
	// jim.print()

	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()

}

// go is a "pass by value" language as opposed to javascript which is "pass by reference"
// this means, when we pass a struct to a function, go is going to clone the value of said
// struct into a new memory in address and use that instead. Think of immutability
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
