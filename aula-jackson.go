package main

import "fmt"

func main() {
	a := 0
	b := 0
	c := 0

	fmt.Scan(&a, &b, &c)

	if (a+b < c) || (c+a < b) || (b+c < a) {
		fmt.Print(" Não é um triãngulo \n ")
	} else {
		if a == b && b == c {
			fmt.Print(" Triângulo Équilatero  \n ")
		} else {
			if (a == b && a != c) || (b == c && b != a) || (c == a && c != b) {
				fmt.Print(" Triãngulo isóceles \n ")
			} else {
				fmt.Print(" Triânglo escaleno \n ")
			}
		}
	}

}
