package main

import "fmt"

func main() {
	x := 13 % 3
	fmt.Println("odd")
	if x == 1 {
		fmt.Print("Odd")
	} else {
		fmt.Println("even")
	}

	for i := 1; i < 70; i++ {
		if i%2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}
}

// invece del risultato del "/" con il remainder ottieni l'avanzo "%"
