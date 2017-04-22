package main

import "fmt"

func main() {
	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	fmt.Println("[1:2]")
	fmt.Println(greeting[1:2])
	fmt.Println("[:2] ")
	fmt.Println(greeting[:2])
	fmt.Println("[:5] ")
	fmt.Println("[:] ")
	fmt.Println(greeting[:])
}
