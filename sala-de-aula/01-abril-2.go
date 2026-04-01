package main

import "fmt"

//MEDIA DE VALORES
func main() {

	var n1, n2, n3 float64
	fmt.Println(" Digite o primeiro número: ")
	fmt.Scan(&n1)
	fmt.Println("Digite o segundo número: ")
	fmt.Scan(&n2)
	fmt.Println(" Digite o terceiro número")
	fmt.Scan(&n3)

	fmt.Printf(" O maior número é %.2f \n", media(n1, n2, n3))

}
func media(n1, n2, n3 float64) float64 {
	var media float64
	media = (n1 + n2 + n3) / 3

	return media

}
