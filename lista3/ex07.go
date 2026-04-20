package main

import (
	"fmt"
)

func main() {
	var num int

	soma := 0
	quantidade := 0
	maior := 0
	menor := 0

	somaPares := 0
	contPares := 0
	contImpares := 0

	for {
		fmt.Print("Digite um número (30000 para encerrar): ")
		fmt.Scan(&num)

		if num == 30000 {
			break
		}

		if quantidade == 0 {
			maior = num
			menor = num
		}

		soma += num
		quantidade++

		if num > maior {
			maior = num
		}
		if num < menor {
			menor = num
		}

		if num%2 == 0 {
			somaPares += num
			contPares++
		} else {
			contImpares++
		}
	}

	fmt.Println("\n--- RESULTADOS ---")

	if quantidade > 0 {
		media := float64(soma) / float64(quantidade)
		fmt.Println("Soma:", soma)
		fmt.Println("Quantidade:", quantidade)
		fmt.Printf("Média: %.2f\n", media)
		fmt.Println("Maior número:", maior)
		fmt.Println("Menor número:", menor)

		if contPares > 0 {
			mediaPares := float64(somaPares) / float64(contPares)
			fmt.Printf("Média dos pares: %.2f\n", mediaPares)
		} else {
			fmt.Println("Não há números pares.")
		}

		percentualImpares := (float64(contImpares) / float64(quantidade)) * 100
		fmt.Printf("Percentual de ímpares: %.2f%%\n", percentualImpares)

	} else {
		fmt.Println("Nenhum número foi digitado.")
	}
}
