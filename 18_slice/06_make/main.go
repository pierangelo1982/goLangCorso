package main

import "fmt"

func main() {
	customerNumber := make([]int, 3)
	// 3 is length & capacity
	// // length - number of elements referred to by the slice
	// // capacity - number of elements in the underlying array
	// 3 è la lunghezza e la capacità
	// // lenght - numero di elementi referenti la slide
	// // capacity - numero di elementi nell'array
	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array
	// you could also do it like this
	// 3 è la lunghezza - numero di elementi referiti dalla slice
	// 5 is capacity - numero di elementi del sotto array
	// tu puoi anche fare cosi
	greeting[0] = "good morning!"
	greeting[1] = "bonjour!"
	greeting[2] = "dias!"

	fmt.Println(greeting[2])
}
