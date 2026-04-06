package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Print("Digite o valor de A: ")
	fmt.Scan(&a)

	fmt.Print("Digite o valor de B: ")
	fmt.Scan(&b)

	fmt.Print("Digite o valor de C: ")
	fmt.Scan(&c)

	// Verifica se é realmente uma equação do 2º grau
	if a == 0 {
		fmt.Println("Não é uma equação do segundo grau!")
		return
	}

	delta := b*b - 4*a*c

	if delta < 0 {
		fmt.Println("RAÍZES IMAGINÁRIAS")
	} else if delta == 0 {
		x := -b / (2 * a)
		fmt.Println("RAIZ ÚNICA")
		fmt.Printf("x = %.2f\n", x)
	} else {
		x1 := (-b + math.Sqrt(delta)) / (2 * a)
		x2 := (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Println("RAÍZES DISTINTAS")
		fmt.Printf("x1 = %.2f\n", x1)
		fmt.Printf("x2 = %.2f\n", x2)
	}
