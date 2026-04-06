package main

import "fmt"

func main() {

	var num float64
	fmt.Println(" Insira um numero: ")
	fmt.Scan(&num)

	if num > 20 && num < 90 {
		fmt.Printf(" O numero %.2f esta entre 20 e 90", num)
	} else {
		fmt.Printf(" O numero %.2f não esta entre 20 e 90", num)
	}

}
