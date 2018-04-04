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

	/*
		alimpk := person{
			firstName: "Ali",
			lastName: "Mohammadpour"
			contactInfo: contactInfo{
				email: "alimpk[at]outlook[dot]com"
				zipCode: 12345
			}
		}
	*/

	var alimpk person

	alimpk.firstName = "Ali"
	alimpk.lastName = "Mohammadpour"
	alimpk.contactInfo.email = "alimpk[at]outlook[dot]com"
	alimpk.contactInfo.zipCode = 12345
	fmt.Println(alimpk)
	fmt.Printf("%v\n", alimpk)
	alimpk.print()

	alimpk.updateName("Dennis", "Ritchee")

	alimpk.print()

}

func (p *person) updateName(firstName string, lastName string) {
	(*p).firstName = firstName
	(*p).lastName = lastName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
