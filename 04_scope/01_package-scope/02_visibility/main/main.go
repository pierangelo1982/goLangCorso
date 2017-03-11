package main

import (
	"fmt"

	"github.com/pierangelo1982/goLangCorso/04_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
