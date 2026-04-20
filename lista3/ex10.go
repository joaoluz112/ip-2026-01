package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Print("Digite a quantidade de termos (N >= 3): ")
	fmt.Scan(&N)

	a := 0
	b := 1

	fmt.Print("Sequência de Fibonacci: ")
	fmt.Print(a, " ", b, " ")

	for i := 3; i <= N; i++ {
		proximo := a + b
		fmt.Print(proximo, " ")

		a = b
		b = proximo
	}
}
