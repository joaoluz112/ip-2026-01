package main

import "fmt"

func main() {

	var n1, n2 int

	fmt.Print(" Digite os números inteiros: \n ")
	fmt.Scan(&n1, &n2)

	if n1%n2 == 0 {
		fmt.Printf(" O número %d é divisível por %d \n", n1, n2)
	} else {
		fmt.Printf(" O número %d não é divisível por %d ", n1, n2)
	}

}
