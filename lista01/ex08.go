package main

import (
	"fmt"
)

func main() {
	var r, h float64
	const pi = 3.14159
	const custoPorM2 = 100.0

	// Entrada
	fmt.Scan(&r)
	fmt.Scan(&h)

	// Cálculos
	areaBase := pi * r * r
	areaLateral := 2 * pi * r * h
	areaTotal := 2*areaBase + areaLateral

	custo := areaTotal * custoPorM2

	// Saída
	fmt.Printf("O VALOR DO CUSTO E = %.2f\n", custo)
}
