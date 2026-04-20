package main

import (
	"fmt"
)

func main() {
	var b, n int

	fmt.Print("Digite o valor de b (base): ")
	fmt.Scan(&b)

	fmt.Print("Digite o valor de n (expoente): ")
	fmt.Scan(&n)

	resultado := 1

	for i := 0; i < n; i++ {
		resultado *= b
	}

	fmt.Printf("%d^%d = %d\n", b, n, resultado)
}
