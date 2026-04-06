package main

import (
	"fmt"
)

func main() {
	var preco float64
	var categoria string
	var dia int
	var valorFinal float64

	fmt.Print("Digite o preço normal do DVD: R$ ")
	fmt.Scan(&preco)

	fmt.Print("Digite a categoria (C - comum, L - lançamento): ")
	fmt.Scan(&categoria)

	fmt.Print("Digite o dia da semana (1=Dom, 2=Seg, ..., 7=Sáb): ")
	fmt.Scan(&dia)

	valorFinal = preco

	
	if dia == 2 || dia == 3 || dia == 5 {
		valorFinal = valorFinal * 0.6 
	}
	if categoria == "L" || categoria == "l" {
		valorFinal = valorFinal * 1.15
	}

	fmt.Printf("Valor final a pagar: R$ %.2f\n", valorFinal)
}
