package main

import "fmt"

func main(){

	var n1, n2, soma int

	fmt.Print( " Digite os números a serem somados: \n ")
	fmt.Scan( &n1, &n2 )

	soma = n1 + n2

	if soma > 20 {
		soma += 8
		fmt.Printf(" %d ", soma)
		
	}else{
		soma -= 5
		fmt.Printf(" %d ", soma)
	}
		
	


}
