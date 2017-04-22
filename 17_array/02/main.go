package main

import "fmt"

func main() {
	var x [58]string

	// parto da 65 perchè è la A nell'alfabeto (tabella unicode ascii)
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i) // i è uguale a 65, quindi meno 65 zero, incomincio a immagazzinare l'alfabeto ell'array dal valore/posizione zero in su nell'array
	}
	fmt.Println(x)     // stampa alfabeto
	fmt.Println(x[42]) // stampa k
}
