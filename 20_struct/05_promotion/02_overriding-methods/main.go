package main

import "fmt"

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func (p person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

func (dz doubleZero) Greeting() {
	fmt.Println("Miss Moneypenny, o good to see you.")
}

func main() {
	p1 := person{
		Name: "Ian Flemming",
		Age:  44,
	}

	p2 := doubleZero{
		person: person{
			Name: "James Bond",
			Age:  20,
		},
		LicenseToKill: true,
	}

	p1.Greeting()
	p2.Greeting()
	p2.person.Greeting() // sovrascrivo p2.Greeting()... vedi codice per capire... se la func passa da person... ()...
}
