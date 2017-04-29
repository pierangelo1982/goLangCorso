package main

import "fmt"

type foo int // costruisco un type int chiamato foo (che però sarà type non int)

func main() {
	var myAge foo // infatti qui lo dichiato come type foo non int
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)
}
