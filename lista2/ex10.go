package main

import (
	"fmt"
)

func main() {
	var destino int
	var tipo int
	var preco float64

	fmt.Println("Destinos:")
	fmt.Println("1 - Região Norte")
	fmt.Println("2 - Região Nordeste")
	fmt.Println("3 - Região Centro-Oeste")
	fmt.Println("4 - Região Sul")
	fmt.Print("Escolha o destino: ")
	fmt.Scan(&destino)

	fmt.Print("Viagem com retorno? (1 - Sim / 2 - Não): ")
	fmt.Scan(&tipo)

	// Validação
	if destino < 1 || destino > 4 {
		fmt.Println("Destino inválido!")
		return
	}

	if tipo != 1 && tipo != 2 {
		fmt.Println("Opção de retorno inválida!")
		return
	}

	// Cálculo do preço
	switch destino {
	case 1: // Norte
		if tipo == 1 {
			preco = 900
		} else {
			preco = 500
		}
	case 2: // Nordeste
		if tipo == 1 {
			preco = 650
		} else {
			preco = 350
		}
	case 3: // Centro-Oeste
		if tipo == 1 {
			preco = 600
		} else {
			preco = 350
		}
	case 4: // Sul
		if tipo == 1 {
			preco = 550
		} else {
			preco = 300
		}
	}

	fmt.Printf("Preço da passagem: R$ %.2f\n", preco)
}
