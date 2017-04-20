package main

import "fmt"

func main() {
	age := 44
	changeMe(age)
	fmt.Println(age) // 44
}

func changeMe(z int) {
	fmt.Println(z)
	z = 24
}

// ritorna 44 e non 24 perche vedi esercizio 02

// when changeMe is called on line 8
// the value 44 is being passed as an argument

// quando changeMe è chiamato sulla linea 7
// il valore 44 ès stato passato come un argomento
