package main

import "fmt"

func main() {
	var x [256]int

	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 0; i < 256; i++ {
		x[i] = i
	}
	// il for i serve solo per l'if che interrompe a 50, si può lasciare libero fino a 256 volendo
	for v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
	}
}
