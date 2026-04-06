package main

import (
	"fmt"
)

func main() {
	var f, polegadas float64

	// Entrada
	fmt.Scan(&f)
	fmt.Scan(&polegadas)

	// Conversões
	celsius := (5*f - 160) / 9
	mm := polegadas * 25.4

	// Saída
	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", celsius)
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f\n", mm)
}
