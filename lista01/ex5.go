package main

import (
	"fmt"
)

func main() {
	var conta int
	var consumo float64
	var tipo string
	var valor float64

	fmt.Scan(&conta, &consumo, &tipo)

	switch tipo {
	case "R":
		valor = 5.0 + consumo*0.05

	case "C":
		if consumo <= 80 {
			valor = 500.0
		} else {
			valor = 500.0 + (consumo-80)*0.25
		}

	case "I":
		if consumo <= 100 {
			valor = 800.0
		} else {
			valor = 800.0 + (consumo-100)*0.04
		}
	}

	// Saída
	fmt.Printf("CONTA = %d\n", conta)
	fmt.Printf("VALOR DA CONTA = %.2f\n", valor)
}
