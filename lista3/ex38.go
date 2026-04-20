package main

import (
	"fmt"
)

func main() {
	var cpf string

	fmt.Print("Digite o CPF (apenas números): ")
	fmt.Scan(&cpf)

	if len(cpf) != 11 {
		fmt.Println("CPF inválido.")
		return
	}

	digitos := make([]int, 11)
	for i := 0; i < 11; i++ {
		digitos[i] = int(cpf[i] - '0')
	}

	soma := 0
	peso := 10

	for i := 0; i < 9; i++ {
		soma += digitos[i] * peso
		peso--
	}

	resto := soma % 11
	dig1 := 0
	if resto >= 2 {
		dig1 = 11 - resto
	}

	soma = 0
	peso = 11

	for i := 0; i < 10; i++ {
		soma += digitos[i] * peso
		peso--
	}

	resto = soma % 11
	dig2 := 0
	if resto >= 2 {
		dig2 = 11 - resto
	}

	if dig1 == digitos[9] && dig2 == digitos[10] {
		fmt.Println("CPF válido.")
	} else {
		fmt.Println("CPF inválido.")
	}
}
