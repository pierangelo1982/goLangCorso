package main

import (
	"fmt"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	counter := 0
	for i := 0; i < b.N; i++ {
		if i%3 == 0 {
			counter += i
		} else if i%5 == 0 {
			counter += i
		}
	}
	fmt.Println(counter)
}

// run this at command:
// go test -bench='.*'

//
// go test -bench='.*' 13_exercise-solutions/08_just-fyi/02_benchMark/bm_test.go
//
