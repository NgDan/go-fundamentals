package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// properties can be assinged in the same order they are declared
	// in the struct definition
	// alex := person{"Alex", "Anderson"}

	// we can also pass them by property name, which is much safer
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// or initializing an empty person with "zero values" and later populate them
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jimparty@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v", jim)
}
