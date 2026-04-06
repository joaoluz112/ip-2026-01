package main

import (
	"fmt"
	"math"
)

func main() {
	var opcao int
	var r, h float64

	fmt.Println("Escolha a figura:")
	fmt.Println("1 - Cone")
	fmt.Println("2 - Cilindro")
	fmt.Println("3 - Esfera")
	fmt.Print("Opção: ")
	fmt.Scan(&opcao)

	switch opcao {
	case 1:
		fmt.Print("Digite o raio: ")
		fmt.Scan(&r)
		fmt.Print("Digite a altura: ")
		fmt.Scan(&h)

		volume := (1.0 / 3.0) * math.Pi * r * r * h
		geratriz := math.Sqrt(r*r + h*h)
		area := math.Pi * r * (r + geratriz)

		fmt.Printf("Volume do cone: %.2f\n", volume)
		fmt.Printf("Área do cone: %.2f\n", area)

	case 2:
		fmt.Print("Digite o raio: ")
		fmt.Scan(&r)
		fmt.Print("Digite a altura: ")
		fmt.Scan(&h)

		volume := math.Pi * r * r * h
		area := 2 * math.Pi * r * (r + h)

		fmt.Printf("Volume do cilindro: %.2f\n", volume)
		fmt.Printf("Área do cilindro: %.2f\n", area)

	case 3:
		fmt.Print("Digite o raio: ")
		fmt.Scan(&r)

		volume := (4.0 / 3.0) * math.Pi * math.Pow(r, 3)
		area := 4 * math.Pi * r * r

		fmt.Printf("Volume da esfera: %.2f\n", volume)
		fmt.Printf("Área da esfera: %.2f\n", area)

	default:
		fmt.Println("Opção inválida!")
	}
}
