package main

import "fmt"

var x int // in automatico acquisisce valore 0 se non dichiarato diversamente
func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment()) //dovrebbe ritornare 1
	fmt.Println(increment()) // dovrebbe ritornarare 2
	// se aggiungessi un fmt.Println(increment()) ritornerebbe 3 via cosi
}
