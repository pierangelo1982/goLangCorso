package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

// il type double zero contiene il type person e LicenseToSkill
// praticamente il type person diventa parte del type doubleZero
type doubleZero struct {
	person
	LicenceToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bong",
			Age:   20,
		},
		LicenceToKill: true,
	}

	p2 := doubleZero{
		person: person{
			First: "Miss",
			Last:  "Moneypenny",
			Age:   18,
		},
		LicenceToKill: false,
	}

	fmt.Println(p1.First, p1.Last, p1.LicenceToKill)
	fmt.Println(p2.First, p2.Last, p2.LicenceToKill)
}
