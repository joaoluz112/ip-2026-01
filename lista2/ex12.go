package main

import (
	"fmt"
)

func main() {
	var idade int

	fmt.Print("Digite a idade: ")
	fmt.Scan(&idade)

	if idade >= 0 && idade <= 2 {
		fmt.Println("Classificação: Recém-nascido")
	} else if idade >= 3 && idade <= 11 {
		fmt.Println("Classificação: Criança")
	} else if idade >= 12 && idade <= 19 {
		fmt.Println("Classificação: Adolescente")
	} else if idade >= 20 && idade <= 55 {
		fmt.Println("Classificação: Adulto")
	} else {
		fmt.Println("Classificação: Idoso")
	}
}
