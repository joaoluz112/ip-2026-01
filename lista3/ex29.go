package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&N)

	if N <= 0 {
		fmt.Println("Valor inválido.")
		return
	}

	soma := 0

	for i := 1; i <= N; i++ {
		soma += i
	}

	fmt.Printf("Somatório de 1 até %d = %d\n", N, soma)
}
