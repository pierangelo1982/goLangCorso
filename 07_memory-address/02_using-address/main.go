package main

import "fmt"

const metersToYard float64 = 1.09361

func main() {
	var meter float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meter)
	yards := meter * metersToYard
	fmt.Println(meter, "meter is", yards, " yards.")
}
