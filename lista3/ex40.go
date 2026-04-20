package main

import (
	"fmt"
)

func main() {
	preco := 6.0
	ingressos := 130
	despesa := 300.0

	lucroMax := -999999.0
	precoMax := 0.0
	ingressosMax := 0

	fmt.Println("Preço\tIngressos\tLucro")

	for preco >= 1.0 {
		receita := preco * float64(ingressos)
		lucro := receita - despesa

		fmt.Printf("R$ %.2f\t%d\t\tR$ %.2f\n", preco, ingressos, lucro)

		if lucro > lucroMax {
			lucroMax = lucro
			precoMax = preco
			ingressosMax = ingressos
		}

		preco -= 0.6
		ingressos += 30
	}

	fmt.Println("\n--- LUCRO MÁXIMO ---")
	fmt.Printf("Lucro máximo: R$ %.2f\n", lucroMax)
	fmt.Printf("Preço do ingresso: R$ %.2f\n", precoMax)
	fmt.Printf("Ingressos vendidos: %d\n", ingressosMax)
}
