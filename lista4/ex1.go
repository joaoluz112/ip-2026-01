package main

import "fmt"

func pow(x int, n int) int {
	if n == 0 {
		return 1
	}

	return x * pow(x, n-1)
}

func main() {

	var x, n int

	fmt.Println(" INSIRA UM NÚMERO INTEIRO COM SEU RESPECTIVO EXPOENTE: ")
	fmt.Scan(&x, &n)

	pow(x, n)

	fmt.Printf(" O resultado é: %d ", pow(x, n))
}
