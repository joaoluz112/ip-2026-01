package main

import (
	"fmt"
)

func main() {
	var n1, n2 int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&n1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&n2)

	resultado := 0

	positivo := true

	if n2 < 0 {
		n2 = -n2
		positivo = !positivo
	}

	if n1 < 0 {
		n1 = -n1
		positivo = !positivo
	}

	for i := 0; i < n2; i++ {
		resultado += n1
	}

	if !positivo {
		resultado = -resultado
	}

	fmt.Printf("Resultado da multiplicação: %d\n", resultado)
}
