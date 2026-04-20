package main

import (
	"fmt"
)

func main() {
	var X float64

	fmt.Print("Digite um valor real para X: ")
	fmt.Scan(&X)

	S := X
	fatorial := 1.0
	sinal := -1.0

	for i := 1; i <= 20; i++ {
		fatorial *= float64(i)

		termo := X / fatorial

		S += sinal * termo

		sinal *= -1
	}

	fmt.Printf("Resultado da série: %.6f\n", S)
}
