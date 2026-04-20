package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Print("Digite um número na base 8: ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("Valor inválido.")
		return
	}

	decimal := 0
	base := 1 

	for num > 0 {
		digito := num % 10

		if digito >= 8 {
			fmt.Println("Número inválido na base 8.")
			return
		}

		decimal += digito * base
		base *= 8
		num /= 10
	}

	fmt.Println("Equivalente em decimal:", decimal)
}
