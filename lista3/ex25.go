package main

import (
	"fmt"
)

func main() {
	S := 0.0
	numerador := 1.0
	sinal := 1.0

	for k := 15; k >= 1; k-- {
		denominador := float64(k * k)

		S += sinal * (numerador / denominador)

		numerador *= 2
		sinal *= -1
	}

	fmt.Printf("Valor de S: %.6f\n", S)
}
