package main

import "fmt"

func main() {

	resultado := 1
	var base, exp int

	fmt.Println(" Digite o número da base: ")
	fmt.Scan(&base)
	fmt.Println(" Digite o expoente da base: ")
	fmt.Scan(&exp)

	for i := 0; i < exp; i++ {
		resultado *= base
	}
	fmt.Printf(" O número %d elevado ao expoente %d é igual à %d \n", base, exp, resultado)

}
