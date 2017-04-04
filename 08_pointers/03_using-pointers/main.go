package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x20818a220

	var b = &a
	fmt.Println(b)  // 0x20818a220
	fmt.Println(*b) // 43

	*b = 42        // b says, "The value at this address, change it to 42"
	fmt.Println(a) // 42
}

// this is useful
// we can pass a memory address instead of a bunch of values (we can pass a reference)
// and then we can still change the value of whatever is stored at that memory address
// this makes our programs more performant
// we don't have to pass around large amounts of data
// we only have to pass around addresses

// everything is PASS BY VALUE in go, btw
// when we pass a memory address, we are passing a value

// questo è utile
// noi possiamo passare un indirizzo di memoria invece di un mucchio di valori (noi possiamo passare la referenza)
// e allora potremo cambiare il valore of qualsiasi cosa è archiviata in quell'indirizzo di memoria
// questo rende il nostro programma più performante
// cosi non dobbiamo passare un grosso ammontare di dati
// ma dobbiamo passarlo solo agli indirizzi

// ogni cossa è passata dal valore in go, btw
// quando passiamo un indirizzo di memoria, noi passiamo un valore
