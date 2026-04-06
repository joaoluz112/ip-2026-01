package main

import (
	"fmt"
)

func main() {
	var preco float64
	var codigo int
	var valorFinal float64

	fmt.Print("Digite o preço do produto: R$ ")
	fmt.Scan(&preco)

	fmt.Println("Escolha a condição de pagamento:")
	fmt.Println("1 - À vista (dinheiro/cheque) - 10% de desconto")
	fmt.Println("2 - À vista (cartão) - 5% de desconto")
	fmt.Println("3 - 2x sem juros")
	fmt.Println("4 - 3x com 10% de juros")
	fmt.Print("Código: ")
	fmt.Scan(&codigo)

	switch codigo {
	case 1:
		valorFinal = preco * 0.9
		fmt.Printf("Valor final: R$ %.2f\n", valorFinal)

	case 2:
		valorFinal = preco * 0.95
		fmt.Printf("Valor final: R$ %.2f\n", valorFinal)

	case 3:
		valorFinal = preco
		parcela := valorFinal / 2
		fmt.Printf("Valor final: R$ %.2f\n", valorFinal)
		fmt.Printf("2 parcelas de: R$ %.2f\n", parcela)

	case 4:
		valorFinal = preco * 1.10
		parcela := valorFinal / 3
		fmt.Printf("Valor final: R$ %.2f\n", valorFinal)
		fmt.Printf("3 parcelas de: R$ %.2f\n", parcela)

	default:
		fmt.Println("Código inválido!")
	}
}
