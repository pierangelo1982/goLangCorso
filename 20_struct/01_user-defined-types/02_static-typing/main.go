package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)

	var yourAge int
	yourAge = 29
	fmt.Printf("%T %v \n", yourAge, yourAge)

	// this doesn't work:
	// non fuziona perchè uno dei due è type foo (anche se ugualmente integer)
	//	 fmt.Println(myAge + yourAge)

	// this conversion works:
	// funziona perchè converto il type foo in int
	//	 fmt.Println(int(myAge) + yourAge)

}
