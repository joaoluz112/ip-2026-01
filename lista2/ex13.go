package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Print("Digite um número inteiro de 3 casas: ")
	fmt.Scan(&num)

	
	if num < 100 || num > 999 {
		fmt.Println("Erro: o número deve ter exatamente 3 casas!")
		return
	}

	dezena := (num / 10) % 10

	fmt.Printf("Algarismo da dezena: %d\n", dezena)
}
