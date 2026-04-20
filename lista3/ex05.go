package main

import (
	"fmt"
)

func main() {
	var idade int
	var altura, peso float64
	var opcao int

	contMais50 := 0
	somaAltura := 0.0
	contAltura := 0
	contPeso40 := 0
	totalPessoas := 0

	for {
		fmt.Print("Digite a idade: ")
		fmt.Scan(&idade)

		fmt.Print("Digite a altura: ")
		fmt.Scan(&altura)

		fmt.Print("Digite o peso: ")
		fmt.Scan(&peso)

		totalPessoas++

		if idade > 50 {
			contMais50++
		}

		if idade >= 10 && idade <= 20 {
			somaAltura += altura
			contAltura++
		}

		if peso < 40 {
			contPeso40++
		}
    
		fmt.Print("Deseja continuar? (1 - Sim / outro valor - Não): ")
		fmt.Scan(&opcao)

		if opcao != 1 {
			break
		}
	}
  
	fmt.Println("\n--- RESULTADOS ---")

	fmt.Println("Pessoas com mais de 50 anos:", contMais50)

	if contAltura > 0 {
		mediaAltura := somaAltura / contAltura
		fmt.Printf("Média das alturas (10 a 20 anos): %.2f\n", mediaAltura)
	} else {
		fmt.Println("Não há pessoas entre 10 e 20 anos para calcular a média.")
	}

	if totalPessoas > 0 {
		percentual := (float64(contPeso40) / float64(totalPessoas)) * 100
		fmt.Printf("Porcentagem com peso < 40kg: %.2f%%\n", percentual)
	}
}
