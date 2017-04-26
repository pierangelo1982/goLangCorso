package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"zero":  "Good Morning!",
		"one":   "Bonjoir!",
		"two":   "Buenos Dias!",
		"three": "Bongiorno!",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, "two")
	fmt.Println(myGreeting)
}
