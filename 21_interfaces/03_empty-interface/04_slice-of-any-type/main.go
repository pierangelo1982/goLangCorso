package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func main() {
	fido := dog{animal{"bau bau"}, true}
	fifi := cat{animal{"miao miao"}, true}
	shadow := dog{animal{"bau bau"}, true}
	critters := []interface{}{fido, fifi, shadow}
	fmt.Println(critters)
}
