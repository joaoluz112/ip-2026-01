package main

import "fmt"

func main() {
	soma := 0
	quantidade := 0

	fmt.Println(" Números pares entre 50 e 70 ")
	fmt.Println(" Soma e Média")

	for i := 50; i < 70; i++ {
		if i%2 == 0 {
			soma += i
			quantidade++
		}

		media := float64(soma) / float64(quantidade)

		fmt.Printf(" Quantidade de pares: %d \n", quantidade)
		fmt.Printf(" Soma dos pares: %d \n", soma)
		fmt.Printf(" Média: %.1f\n", media)

	}

}
