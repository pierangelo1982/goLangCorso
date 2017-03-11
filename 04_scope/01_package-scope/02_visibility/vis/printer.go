package vis

import "fmt"

// PrintVar is exported because start with capital letter
// PrintVar è esportabile perchè inzia cn lettera maiuscola
func PrintVar() {
	fmt.Println(MyName)
	fmt.Println(yourName) //yourname lo vede anche se minuscolo solo xchè dentro lo stesso pacchetto
}
