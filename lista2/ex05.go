package main

import "fmt"

func main(){

	var n int

	fmt.Print(" Digite um número inteiro: \n ")
	fmt.Scan( &n )

	if n % 5 == 0 {
		fmt.Print(" O número é divisível por 5 \n")	
	}else{
		fmt.Print(" O número não é divisível por 5")
	}

}
