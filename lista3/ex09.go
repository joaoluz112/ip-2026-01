package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Print("Digite a quantidade de alunos: ")
	fmt.Scan(&N)

	var nota1, nota2 float64
	somaMedias := 0.0

	aprovados := 0
	exame := 0
	reprovados := 0

	for i := 1; i <= N; i++ {
		fmt.Printf("\nAluno %d\n", i)

		fmt.Print("Digite a primeira nota: ")
		fmt.Scan(&nota1)

		fmt.Print("Digite a segunda nota: ")
		fmt.Scan(&nota2)

		media := (nota1 + nota2) / 2
		somaMedias += media

		fmt.Printf("Média: %.2f - ", media)

		if media <= 3 {
			fmt.Println("Reprovado")
			reprovados++
		} else if media < 7 {
			fmt.Println("Exame")
			exame++
		} else {
			fmt.Println("Aprovado")
			aprovados++
		}
	}

	mediaClasse := somaMedias / float64(N)

	fmt.Println("\n--- RESULTADOS ---")
	fmt.Println("Total de aprovados:", aprovados)
	fmt.Println("Total em exame:", exame)
	fmt.Println("Total de reprovados:", reprovados)
	fmt.Printf("Média da classe: %.2f\n", mediaClasse)
}
