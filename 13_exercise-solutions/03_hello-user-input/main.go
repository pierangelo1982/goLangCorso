package main

import "fmt"

func main() {
	var name string
	fmt.Print("Inserisci il tuo Nome: ")
	fmt.Scan(&name)
	fmt.Println("Ciao", name)
}
