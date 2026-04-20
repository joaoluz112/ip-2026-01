package main

import (
	"fmt"
)

func main() {
	H := 0.0
	numerador := 1.0

	for i := 1; i <= 50; i++ {
		H += numerador / float64(i)
		numerador += 2
	}

	fmt.Printf("Valor de H: %.2f\n", H)
}
