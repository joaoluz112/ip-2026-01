package main

import (
	"fmt"
)

func main() {
	var x float64

	fmt.Print("Digite o valor de x: ")
	fmt.Scan(&x)

	if x == 2 {
		fmt.Println("Erro: divisão por zero (x não pode ser 2)")
		return
	}

	fx := 8 / (2 - x)

	fmt.Printf("f(%.2f) = %.2f\n", x, fx)
}
