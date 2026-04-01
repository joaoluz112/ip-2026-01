package main

import "fmt"

func main(){

// ENTRADA

    var A, B, C, D int
	var vcar float64
	fmt.Println(" Qual o valor do carro: ")
	fmt.Scan( &vcar )
	fmt.Println(" Digite 1 para SIM e 0 para NÃO")
	fmt.Println(" Complementos adicioanis adicionais: ")
	fmt.Println("Ar condicionado: ")
	fmt.Scan(&A)
	fmt.Println("Pintura metálica: ")
	fmt.Scan(&B)
	fmt.Println("Vidro elétrico: ")
	fmt.Scan(&C)
	fmt.Println("Direção hidráulica: ")
	fmt.Scan(&D)
    
//CONDIÇÕES

	if A == 1 {
		vcar += 1750	
	}
	if B == 1{
		vcar += 800
	}
	if C == 1{
		vcar += 1200
	}
	if D == 1{
		vcar += 2000
	}

//SAÍDA

	fmt.Printf(" O valor total do carro será de R$ %.2f \n", vcar)


 
}
