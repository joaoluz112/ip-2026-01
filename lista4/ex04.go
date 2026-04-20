package main

import (
	"fmt"
)

func decimalParaBinario(n int) string {
	if n == 0 {
		return "0"
	}
	if n == 1 {
		return "1"
	}

	return decimalParaBinario(n/2) + fmt.Sprintf("%d", n%2)
}

func main() {
	var num int

	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("Valor inválido.")
		return
	}

	resultado := decimalParaBinario(num)

	fmt.Println("Binário:", resultado)
}
