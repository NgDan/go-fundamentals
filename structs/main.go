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

	// the & in front of the variable means we're accessing
	// the memory address of the value
	jimPointer := &jim
	fmt.Println(jimPointer)
	jimPointer.updateName("jimmy")
	jim.print()

}

// go is a "pass by value" language as opposed to javascript which is "pass by reference"
// this means, when we pass a struct to a function, go is going to clone the value of said
// struct into a new memory in address and use that instead. Think of immutability
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

// the * in front of the variable means we're accessing the
// value this memory address points to. Since the receiver here
// is a pointer (jimPointer), we need to dereference it by using
// *pointerToPerson in order to get the value at that address.
// Also, *person is a TYPE, not an actual dereferenced value.
// It only means that the pointerToPerson has to be of type
// pointer to person
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
