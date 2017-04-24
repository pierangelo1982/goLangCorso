package main

import "fmt"

func main() {
	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

	mySlice = append(mySlice[:2], mySlice[3:]...) // arriva fino al 2 e riparti da dopo il 3, cosi cancella quello in mezzo rifacendo l'append
	fmt.Println(mySlice)
}

// arriva fino al 2 e riparti da dopo il 3,
// cosi cancella quello in mezzo rifacendo l'append
