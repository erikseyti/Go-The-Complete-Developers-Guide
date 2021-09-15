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
	// First Way to create a person
	// alex := person{"Alex", "Anderson"}

	// Second way to create a person
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// Third Way to create a struct person
	// var alex person

	// fmt.Println(alex)
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// alex.contact.email = "alex@hotmail.net"
	// alex.contact.zipCode = 5111

	// Print exibe todos os valores da struct Person
	// fmt.Printf("%+v", alex)
	// fmt.Println("")

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "Jim@gamil.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v", jim)
	fmt.Println("")

}
