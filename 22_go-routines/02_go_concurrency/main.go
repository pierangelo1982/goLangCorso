package main

import "fmt"

func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ", i)
	}
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar: ", i)
	}
}

// funziona ma non stampa, per il semplice motivo che le 2 funzioni avvengono contemporaneamente
