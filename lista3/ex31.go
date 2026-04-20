package main

import (
	"fmt"
)

func main() {
	var graos uint64 = 1
	var total uint64 = 0

	for i := 1; i <= 64; i++ {
		total += graos

		fmt.Printf("Casa %d: %d grãos\n", i, graos)

		graos *= 2
	}

	fmt.Println("\nTotal de grãos:", total)
}
