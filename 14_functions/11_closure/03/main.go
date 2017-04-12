package main

import "fmt"

func main() {
	x := 0 // quindi sarebbe lo stesso fare var x int;
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	// se aggiungessi un fmt.Println(increment()) ritornerebbe 3 via cosi
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
anonymous function
a function without a name
func expression
assigning a func to a variable
*/

/*
  le graffe ci aiutano a limitare gli scopi di variabili usate da funzioni multiple,
  senza graffe, per due o più funzioni avrebbero accesso alla stessa variabile,
  questa variabile avrebbe bisogno di essere isnerita in un package scope
  funzione anonima
  una funzione senza nome
  funzione espressione
  assegna una funzione ad una variabile
*/
