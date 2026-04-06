package main

import (
	"fmt"
)

func main() {
	var tipo string
	var consumo float64
	var valor float64

	fmt.Print("Digite o tipo de consumidor (R - Residencial, C - Comercial, I - Industrial): ")
	fmt.Scan(&tipo)

	fmt.Print("Digite o consumo de água (em m³): ")
	fmt.Scan(&consumo)

	switch tipo {
	case "R", "r":
		valor = 5.0 + (consumo * 0.05)

	case "C", "c":
		if consumo <= 80 {
			valor = 500.0
		} else {
			valor = 500.0 + (consumo-80)*0.25
		}

	case "I", "i":
		if consumo <= 100 {
			valor = 800.0
		} else {
			valor = 800.0 + (consumo-100)*0.04
		}

	default:
		fmt.Println("Tipo de consumidor inválido!")
		return
	}

	fmt.Printf("Valor da conta: R$ %.2f\n", valor)
}
