package main

import "fmt"

func main() {
	// Set up the pipeline and consume the output.
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n) // 16 then 81
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// il calcolo che fa per ottenere 16:
//2 * 2 = 4; 4 * 4 = 16
// il calocolo che fa x ottenere 81
// 3 * 3 = 9; 9 * 9 = 81
// questo perchè c'è la funzione inclusa nella stessa funzione sq:
// range sq(sq(gen(2, 3))) {
