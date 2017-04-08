package main

import "fmt"

func main() {
	counter := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			counter += i
		} else if i%5 == 0 {
			counter += i
		}
	}

	fmt.Println(counter)
}

/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.
*/

/*
se noi listiamo tutti i numero naturali sotto 10 che sono multipli di 3 o 5, noi otteniamo 3, 5, 6 e 9.
la somma di questi multipli è 23.
Somma tutti i multipli di 3 o 5 sotto 1000
*/
