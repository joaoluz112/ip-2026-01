package main

import "fmt"

func main(){

	var n int

	fmt.Print( " Digite um número: \n" )
	fmt.Scan( &n )

	if ( n < 0) {
		fmt.Print( " O número é negativo \n " )		
	}else if n > 0{
		fmt.Print( " O número é positivo \n ")
	}else{
		fmt.Print( " O número é nulo \n ")
	}
}
