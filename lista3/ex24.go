package main

import (
	"fmt"
)

func fatorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	var A float64

	fmt.Print("Digite o ângulo em radianos: ")
	fmt.Scan(&A)

	seno := 0.0
	sinal := 1.0

	for i := 1; i <= 7; i += 2 {
		potencia := 1.0
		for j := 1; j <= i; j++ {
			potencia *= A
		}

		termo := potencia / float64(fatorial(i))

		seno += sinal * termo

		sinal *= -1
	}

	fmt.Printf("sen(%.2f) ≈ %.6f\n", A, seno)
}
