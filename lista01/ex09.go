package main

import (
	"fmt"
)

func main() {
	var A, B, C float64

	// Entrada (cada valor em uma linha)
	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)

	// Cálculo do delta
	delta := B*B - 4*A*C

	// Saída
	fmt.Printf("O VALOR DE DELTA E = %.2f\n", delta)
}
