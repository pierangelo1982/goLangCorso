package main

import "fmt"

func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/

/*
parentesi ci aiutano a limitare gli scope di variabili usate da multiple funzioni
senza parentesi, per due o più funzioni avrebbero accesso alla stessa variabile,
questa variabile avrebbe bisogno di essere un package scope
*/
