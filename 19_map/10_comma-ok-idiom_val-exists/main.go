package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(myGreeting)

	//delete(myGreeting, 2)

	if val, exists := myGreeting[2]; exists {
		fmt.Println("That Value exists.")
		fmt.Println("val ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exists.")
		fmt.Println("val ", val)
		fmt.Println("exists: ", exists)
	}
}
