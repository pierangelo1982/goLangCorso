package main

import "fmt"

var a = "questo archivia la variabile a"
var b, c string = "archivia in b", "archivia in c"
var d string

func main() {
	d = "archivia in d" // dichiarata sopra
	var e = 42
	f := 43
	g := "archivia in g"
	h, i := "archivia in h", "archivia in i"
	j, k, l, m := 44.7, true, false, "archivia in m"
	n := "archivia in n"
	o := "archivia in o"

	fmt.Println("a -", a)
	fmt.Println("b -", b)
	fmt.Println("c -", c)
	fmt.Println("d -", d)
	fmt.Println("e -", e)
	fmt.Println("f -", f)
	fmt.Println("g -", g)
	fmt.Println("h -", h)
	fmt.Println("i -", i)
	fmt.Println("j -", j)
	fmt.Println("k -", k)
	fmt.Println("l -", l)
	fmt.Println("m -", m)
	fmt.Println("n -", n)
	fmt.Println("o -", o)
}
