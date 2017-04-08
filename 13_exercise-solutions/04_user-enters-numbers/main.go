package main

import "fmt"

func main() {
	var bigNumber int
	var smallNumber int

	fmt.Print("Inserisci il primo numero: ")
	fmt.Scan(&bigNumber)

	fmt.Print("Inserisci un secondo numero inferiore al primo: ")
	fmt.Scan(&smallNumber)

	fmt.Println(bigNumber, "%", smallNumber, "=", bigNumber/smallNumber)
}
