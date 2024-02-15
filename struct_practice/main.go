package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	contact   contactInfo // field name can be ignored
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// auotmatically assign zero value
	var alex Person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contact.email = "alex.anderson@gmail.com"
	alex.contact.zipCode = 24232

	fmt.Println(alex)
	fmt.Printf("%+v\n", alex) // %+v print all different field name
	alex.print()

	jim := Person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim.party@gmail.com",
			zipCode: 94000,
		},
	}

	jim.print()

	jim.udpateName("Eddie")
	jim.print()

	// Because slice is reference type, so we can pass it into function directly
	// struct is value type data struct, we have to use pointer
	mySlice := []string{"hi", "there", "how", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *Person) udpateName(newFirstName string) {
	(*p).firstName = newFirstName
}
