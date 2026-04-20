package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Valor inválido! Fatorial não existe para números negativos.")
		return
	}

	fatorial := 1

	for i := 1; i <= n; i++ {
		fatorial *= i
	}

	fmt.Printf("%d! = %d\n", n, fatorial)
}
