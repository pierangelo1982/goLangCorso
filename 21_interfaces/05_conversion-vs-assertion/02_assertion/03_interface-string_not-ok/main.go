package main

import "fmt"

func main() {
	var name interface{} = 7
	str, ok := name.(tring)
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
