package main

import "fmt"

func max(x int) int {
	return 42 + x
}

func main() {
	max := max(7)
	fmt.Println(max) // max is now the result and not the function
}

// don't do this; bad coding practice to shadow the variables

// da non fare, pessima pratica oscurare le variabili

// in questo modo la funzione diventa un risultato, perdendo le sue capacità.... pessimo metodo
