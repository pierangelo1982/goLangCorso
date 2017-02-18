package main

import "fmt"

func main() {
	for i := 0; i < 1000000; i++ {
		fmt.Printf("%d \t %b \t %#X \t %q \n", i, i, i, i)
	}
}

// %q è UTF
