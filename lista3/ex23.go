package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Print("Digite o número de termos (N > 0): ")
	fmt.Scan(&N)

	if N <= 0 {
		fmt.Println("Valor inválido.")
		return
	}

	S := 0.0
	numerador := 1000.0
	sinal := 1.0

	for i := 1; i <= N; i++ {
		S += sinal * (numerador / float64(i))

		numerador -= 3
		sinal *= -1
	}

	fmt.Printf("Resultado da série: %.2f\n", S)
}
