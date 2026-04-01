package main

import "fmt"
//MAIOR ENTRE TRẼS NÚMEROS

func main() {

	var n1, n2, n3 int
	fmt.Println(" Digite o primeiro número: ")
	fmt.Scan(&n1)
	fmt.Println("Digite o segundo número: ")
	fmt.Scan(&n2)
	fmt.Println(" Digite o terceiro número")
	fmt.Scan(&n3)

	fmt.Printf(" O maior número é %d \n", maior(n1, n2, n3))

}
func maior(n1, n2, n3 int) int {
	maior := n1

	if n2 > maior {
		maior = n2
	}
	if n3 > maior {
		maior = n3
	}

	return maior

}
