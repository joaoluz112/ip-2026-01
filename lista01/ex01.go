package main

import "fmt"

func main() {
	var n1, n2, n3, media float64

	fmt.Print(" Notas do aluno: \n ")
	fmt.Scan(&n1, &n2, &n3)
	media = (n1 + n2 + n3) / 3
	if media < 6 {
		fmt.Printf(" Média = %.2f ", media)
		fmt.Print(" REPROVADO \n ")

	} else {
		fmt.Printf(" Média = %.2f ", media)
		fmt.Print(" APROVADO \n ")
	}

}
