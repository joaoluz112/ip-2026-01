package main

import (
	"fmt"
)

func main() {
	S := 0.0
	num := 37.0

	for i := 1; i <= 37; i++ {
		S += (num * (num + 1)) / float64(i)
		num--
	}

	fmt.Printf("Valor de S: %.2f\n", S)
}
