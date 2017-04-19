package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1) // scala tutti i numeri fino all'uno sopra dell'if 3,2,1
}

func main() {
	fmt.Println(factorial(4)) // 24
}

// 4 * 3 = 12
// 12 * 2 = 24
// 24 * 1 = 24
