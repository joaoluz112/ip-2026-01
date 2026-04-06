package main

import (
	"fmt"
)

func main() {
	var id int
	var nota1, nota2, nota3, mediaEx float64
	var mediaFinal float64
	var conceito string

	fmt.Print("Digite o número do aluno: ")
	fmt.Scan(&id)

	fmt.Print("Digite a nota 1: ")
	fmt.Scan(&nota1)

	fmt.Print("Digite a nota 2: ")
	fmt.Scan(&nota2)

	fmt.Print("Digite a nota 3: ")
	fmt.Scan(&nota3)

	fmt.Print("Digite a média dos exercícios: ")
	fmt.Scan(&mediaEx)

	mediaFinal = (nota1 + nota2*2 + nota3*3 + mediaEx) / 7
  
	if mediaFinal >= 9.0 {
		conceito = "A"
	} else if mediaFinal >= 7.5 {
		conceito = "B"
	} else if mediaFinal >= 6.0 {
		conceito = "C"
	} else if mediaFinal >= 4.0 {
		conceito = "D"
	} else {
		conceito = "E"
	}

	
	fmt.Println("\n--- Resultado ---")
	fmt.Printf("Aluno: %d\n", id)
	fmt.Printf("Notas: %.2f, %.2f, %.2f\n", nota1, nota2, nota3)
	fmt.Printf("Média dos exercícios: %.2f\n", mediaEx)
	fmt.Printf("Média final: %.2f\n", mediaFinal)
	fmt.Printf("Conceito: %s\n", conceito)

	
	if conceito == "A" || conceito == "B" || conceito == "C" {
		fmt.Println("Situação: APROVADO")
	} else {
		fmt.Println("Situação: REPROVADO")
	}
}
