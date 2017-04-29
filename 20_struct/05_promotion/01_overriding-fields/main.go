package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person
	First         string
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "James", //First string of person
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double Zero Seven", // First string di doubleZero
		LicenseToKill: true,
	}

	p2 := doubleZero{
		person: person{
			First: "Miss", // first of person
			Last:  "Moneypenny",
			Age:   18,
		},
		First:         "If looks could kill",
		LicenseToKill: false,
	}

	// fields and methods of the inner-type are promoted to the outer-type
	fmt.Println(p1.First, p1.person.First) // il secondo First lo chiamo passando da person (James)
	fmt.Println(p2.First, p2.person.First) // // il secondo First lo chiamo passando da person (Miss)

}
