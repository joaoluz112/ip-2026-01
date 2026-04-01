package main

import "fmt"
// O FATORIAL DE UM NÚMERO
func main() {

	var n1 int

	fmt.Println(" Digite um número: ")
	fmt.Scan(&n1)

	fmt.Printf(" O fatorial deste número é %d \n", fatorial(n1))

}
func fatorial(numero int) int {
	for i := numero - 1; i > 0; i-- {
		numero *= i
	}

	return numero

}
