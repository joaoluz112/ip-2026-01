package main

import (
	"fmt"
)

func main() {
	S := 0.0
	numerador := 100.0
	fatorial := 1.0

	for i := 0; i < 20; i++ {
	
		if i > 0 {
			fatorial *= float64(i)
		}

		S += numerador / fatorial

		numerador--
	}

	fmt.Printf("Soma dos 20 termos: %.6f\n", S)
}
