package main

import (
	"fmt"
)

func main() {
	var num int

	for {
		fmt.Print("Digite um número (<= 0 para sair): ")
		fmt.Scan(&num)

		if num <= 0 {
			fmt.Println("Programa encerrado.")
			break
		}

		ehPerfeito := false

		for i := 1; i*i <= num; i++ {
			if i*i == num {
				ehPerfeito = true
				break
			}
		}

		if ehPerfeito {
			fmt.Println(num, "é um quadrado perfeito.")
		} else {
			fmt.Println(num, "não é um quadrado perfeito.")
		}
	}
}
