package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("Valor inválido.")
		return
	}

	if num == 0 {
		fmt.Println("Binário: 0")
		return
	}

	binario := ""

	for num > 0 {
		resto := num % 2
		binario = fmt.Sprintf("%d%s", resto, binario)
		num /= 2
	}

	fmt.Println("Binário:", binario)
}
