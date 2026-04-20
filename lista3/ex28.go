package main

import (
	"fmt"
	"math"
)

func main() {
	S := 0.0
	sinal := 1.0
	num := 1.0

	for i := 0; i < 51; i++ {
		S += sinal * (1.0 / (num * num * num))

		num += 2    
		sinal *= -1  
	}

	pi := math.Sqrt(32 * S)

	fmt.Printf("Valor aproximado de pi: %.10f\n", pi)
}
