package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Print("Digite o valor de N: ")
	fmt.Scan(&N)

	if N <= 0 {
		fmt.Println("Valor inválido.")
		return
	}

	fmt.Println("Série:")

	for i := 1; i <= N; i++ {
		fmt.Print(i*i, " ")
	}
}
