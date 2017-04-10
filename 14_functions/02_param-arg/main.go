package main

import "fmt"

func main() {
	greet("Jane")
	greet("John")
}

func greet(name string) {
	fmt.Println(name)
}

// greet is declared with a parameter
// when calling greet, pass in an argument

// greet è dichiarato con un parametro
// quando chiami greet, gli passi un argomento
