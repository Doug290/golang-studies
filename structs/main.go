package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	jim := person {
		firstName: "Jim",
		lastName: "Pasrty",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 940039,
		},
	}
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}