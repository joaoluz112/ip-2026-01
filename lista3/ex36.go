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
		fmt.Println("Hexadecimal: 0")
		return
	}

	hex := ""
	digitos := "0123456789ABCDEF"

	for num > 0 {
		resto := num % 16
		hex = string(digitos[resto]) + hex
		num /= 16
	}

	fmt.Println("Hexadecimal:", hex)
}
