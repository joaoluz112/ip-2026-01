package main

import (
	"fmt"
	"math"
)

func main() {
	var h, a float64
	var Ab, V float64

	fmt.Scan(&h, &a)

	// Área do hexágono
	Ab = (3 * math.Sqrt(3) / 2) * (a * a)

	// Volume da pirâmide
	V = (1.0 / 3.0) * Ab * h

	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", V)
}
