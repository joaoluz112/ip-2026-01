package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3 int

	fmt.Scan(&n1, &n2, &n3)

	// Validação: cada número deve ter apenas 1 dígito
	if n1 < 0 || n1 > 9 || n2 < 0 || n2 > 9 || n3 < 0 || n3 > 9 {
		fmt.Println("DIGITO INVALIDO")
		return
	}

	// Monta o número
	numero := n1*100 + n2*10 + n3

	// Calcula o quadrado
	quadrado := numero * numero

	// Saída
	fmt.Printf("%d, %d\n", numero, quadrado)
}
