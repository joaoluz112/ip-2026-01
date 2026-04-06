package main

import (
	"fmt"
	"math"
)

func main() {

	var n float64

	fmt.Print(" Digite um número: \n ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Printf(" O seu número ao quadrado é = %f \n ", n*n)

	} else {
		fmt.Printf("A raiz quadrada do seu número é = %.2f \n", math.Sqrt(n))
	}

}
