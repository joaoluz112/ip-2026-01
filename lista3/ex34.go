package main

import (
	"fmt"
)

func main() {
	var n1, n2 int

	fmt.Print("Digite N1: ")
	fmt.Scan(&n1)

	fmt.Print("Digite N2: ")
	fmt.Scan(&n2)

	if n1 <= 0 || n2 <= 0 {
		fmt.Println("Valores inválidos.")
		return
	}

	mmc := n1
	if n2 > n1 {
		mmc = n2
	}

	for {
		if mmc%n1 == 0 && mmc%n2 == 0 {
			break
		}
		mmc++
	}

	fmt.Printf("MMC(%d, %d) = %d\n", n1, n2, mmc)
}
