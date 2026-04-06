package main

import (
	"fmt"
)

func main() {
	var a1, r, n int
	var soma int

	fmt.Scan(&a1, &r, &n)

	termo := a1

	for i := 0; i < n; i++ {
		soma += termo
		termo += r
	}

	fmt.Println(soma)
}
