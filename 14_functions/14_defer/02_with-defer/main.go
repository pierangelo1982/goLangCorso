package main

import "fmt"

func hello() {
	fmt.Print("Hello ")
}

func world() {
	fmt.Println("World")
}

func main() {
	defer world() //deferisco, rimando, posticipo la chiamata di questa funzione
	hello()
}

// altrimenti sarebbe stato world hello

// viene utile quando apri un file che poi deve essere chiuso o salvato...
// puoi farlo subito e posticipare l'azione di salvataggio senza dover andare
// in fondo alla riga dopo tutto codice
