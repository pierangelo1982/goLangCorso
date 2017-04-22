package main

import "fmt"

func main() {
	var x [58]int // dico che l'array deve essere lungo 58 elementi
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42]) // chiedo di mostrarmi cosa c'è nella posizione 42 (riuta zero perchè vuoto)
	x[42] = 777        // memorizzo un valore nell'array alla posizione 42
	fmt.Println(x[42])
}
