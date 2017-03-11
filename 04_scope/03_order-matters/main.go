package main

import "fmt"

func main() {
	fmt.Println(x)
	fmt.Println(y)
	x := 42
}

var y = 42

// la y la vede perchè esterna alla funzione, anche se dichiarata dopo.
// la x no perchè pur essendo dentro la funzione non è in ordine...
// l'ordine conta
