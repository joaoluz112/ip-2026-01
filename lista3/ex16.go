package main

import (
	"fmt"
)

func main() {
	var n int
	var a1, a2 int

	fmt.Print("Digite o primeiro termo: ")
	fmt.Scan(&a1)

	fmt.Print("Digite o segundo termo: ")
	fmt.Scan(&a2)

	fmt.Print("Digite a quantidade de termos (N >= 3): ")
	fmt.Scan(&n)

	if n < 3 {
		fmt.Println("Valor inválido.")
		return
	}

	fmt.Print("Série de Fetuccine: ")

	fmt.Print(a1, " ", a2, " ")

	for i := 3; i <= n; i++ {
		var proximo int

		if i%2 != 0 {
			proximo = a2 + a1
		} else {
			proximo = a2 - a1
		}

		fmt.Print(proximo, " ")

		a1 = a2
		a2 = proximo
	}
}
