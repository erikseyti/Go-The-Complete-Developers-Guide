package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// First Way to create a person
	// alex := person{"Alex", "Anderson"}

	// Second way to create a person
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// Third Way to create a struct person
	var alex person

	fmt.Println(alex)
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	// Print exibe todos os valores da struct Person
	fmt.Printf("%+v", alex)
	fmt.Println("")

}
