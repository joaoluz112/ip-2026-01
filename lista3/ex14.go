package main

import (
	"fmt"
)

func main() {
	var n1, n2 int

	fmt.Print("Digite o valor de N1: ")
	fmt.Scan(&n1)

	fmt.Print("Digite o valor de N2: ")
	fmt.Scan(&n2)

	if n1 > n2 {
		n1, n2 = n2, n1
	}

	fmt.Println("Números primos no intervalo:")

	for i := n1; i <= n2; i++ {
		if ehPrimo(i) {
			fmt.Print(i, " ")
		}
	}
}

func ehPrimo(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
