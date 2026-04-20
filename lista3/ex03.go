package main

import (
	"fmt"
)

func main() {
	var salarioCarlos float64

	fmt.Print("Digite o salário de Carlos: ")
	fmt.Scan(&salarioCarlos)

	salarioJoao := salarioCarlos / 3

	valorCarlos := salarioCarlos
	valorJoao := salarioJoao

	meses := 0

	for valorJoao < valorCarlos {
		valorCarlos *= 1.02 
		valorJoao *= 1.05 
		meses++
	}

	fmt.Println("Meses necessários:", meses)
	fmt.Printf("Valor final de Carlos: %.2f\n", valorCarlos)
	fmt.Printf("Valor final de João: %.2f\n", valorJoao)
}
