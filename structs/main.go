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
	// *pointerToPerson is an actual dereference and will result in the value
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// there's some syntactic sugar in go where if we have a receiver of type pointer of any variable
// we can pass that variable to the receiver without first converting it to a pointer and go will
// automatically get the pointer of that variable for us
// For example here, we do &jim to get the pointer of jim:
// jimPointer := &jim
// jimPointer.updateName("jimmy")

// func (pointerToPerson *person) updateName(newFirstName string) {
// 	(*pointerToPerson).firstName = newFirstName
// }
// But we could also do this:
// jim.updateName("jimmy")
// func (pointerToPerson *person) updateName(newFirstName string) {
// 	(*pointerToPerson).firstName = newFirstName
// }
// here, there receiver is of type person, not *person but go
// will understand and go and get the actual value of person
// instead of cloning person into a different memory address
// and pass it by value
