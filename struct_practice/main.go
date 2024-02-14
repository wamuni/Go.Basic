package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	contact   contactInfo
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

	jim := Person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim.party@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v", jim)

}
