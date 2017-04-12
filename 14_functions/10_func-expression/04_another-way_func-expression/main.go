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
}

// una funzione che ritorna una funzione in questo caso stringa
// funzione makeGreeter dichiara di volere come ritorno una funzione stringa func() string {
// infatti all'interno returna una return func() string {
//  return "Hello World"
//}
