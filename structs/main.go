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
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jim.updateName("jimmy")
	jim.print()

}

//*person is a type description that means it can only receive of a type person
func (pointerToPerson *person) updateName(newFirstName string) {
	//manipulate the value the pointer is referencing
	//take this memory address and point it to the person and turn it on the actual valuey
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
