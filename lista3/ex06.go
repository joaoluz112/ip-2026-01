package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&num)

	if num <= 0 {
		fmt.Println("Número inválido.")
		return
	}

	ehTriangular := false

	for i := 1; i*(i+1)*(i+2) <= num; i++ {
		if i*(i+1)*(i+2) == num {
			fmt.Printf("%d = %d x %d x %d\n", num, i, i+1, i+2)
			ehTriangular = true
			break
		}
	}

	if ehTriangular {
		fmt.Println("É um número triangular.")
	} else {
		fmt.Println("Não é um número triangular.")
	}
}
