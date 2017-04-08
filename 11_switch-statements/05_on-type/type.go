package main

import "fmt"

//  switch on types
//  -- normally we switch on value of variable
//  -- go allows you to switch on type of variable

// switch types
// -- normalmente noi switchiamo sul valore della variabile
// -- go ci permette di switchare sul tipo di variabile
type contact struct {
	greeting string
	name     string
}

// SwitchOnType works with interfaces
// we'll learn more about interfaces later
// switchOnType funziona con l'interfaccia
// noi impareremo più tardi cosa sono
func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting, "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknow")
	}
}

// rihiamando la funzione sopra SwitchOnType, riconosce il formato dei valori presenti e lo stampa
func main() {
	SwitchOnType(7)
	SwitchOnType("McLeod")
	var t = contact{"Good to see you", "Tim"} // praticamente un arrai del type contact creato sora
	SwitchOnType(t)
	SwitchOnType(t.greeting) //Good to see you
	SwitchOnType(t.name)     // Tim
}
