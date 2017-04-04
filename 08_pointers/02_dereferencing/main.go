package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) //// 0x20818a220

	var b = &a
	fmt.Println(b)  // 0x20818a220
	fmt.Println(*b) // 43
}

// b is an int pointer;
// b points to the memory address where an int is stored
// to see the value in that memory address, add a * in front of b
// this is known as dereferencing
// the * is an operator in this case

// b è un int pointer
// b punta all' indirizzo di memoria dove un int è archiviato
// per vedere il valore in questo indirizzo di memoria, aggiungi un * prima di b
// questo procedimento è conosciuto come dereferencing (derifenzazione)
// l' * (asterisco) è un operatore in questo caso
