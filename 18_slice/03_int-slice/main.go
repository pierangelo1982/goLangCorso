package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 3) // 0 è il valore, 3 la capienza dell'array, che poi si autoicrementa per se stesso quando pieno nel loop

	fmt.Println("--------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("--------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "value", mySlice[i])
	}
}
