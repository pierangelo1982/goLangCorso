package main

import "fmt"

func makeGreeter() func() string {
	return func() string { //nota il return alla funzione che a sua volta ha un return
		return "Hello World"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Printf("%T\n", greet)
}
