package main

import (
	"fmt"
)

func main() {
	var salarioMinimo, consumo float64

	fmt.Scan(&salarioMinimo, &consumo)

	// Cálculo do valor de 1 kW
	valorKW := (0.7 * salarioMinimo) / 100

	// Custo total
	custo := consumo * valorKW

	// Custo com desconto de 10%
	custoDesconto := custo * 0.9

	// Saída
	fmt.Printf("Custo por kW: R$ %.2f\n", valorKW)
	fmt.Printf("Custo do consumo: R$ %.2f\n", custo)
	fmt.Printf("Custo com desconto: R$ %.2f\n", custoDesconto)
}
