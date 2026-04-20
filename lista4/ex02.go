package main

import (
	"fmt"
)

func soma(arr []float64) float64 {
  
	if len(arr) == 0 {
		return 0
	}

	return arr[0] + soma(arr[1:])
}

func main() {
	valores := []float64{1.5, 2.0, 3.5, 4.0}

	resultado := soma(valores)

	fmt.Printf("Soma dos valores: %.2f\n", resultado)
}
