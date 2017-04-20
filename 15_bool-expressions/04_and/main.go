package main

import "fmt"

func main() {
	if true && false {
		fmt.Println("this did not run")
	}
}
