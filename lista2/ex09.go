package main

import "fmt"

func main(){

	var vc, vv float64
	fmt.Println( " O valor da compra: ")
	fmt.Scan( &vc )

	switch{
	case vc < 10:
		vv = vc + (vc * 0.7)
		fmt.Printf( " Valor de venda será de R$ %f", vv)

	case vc >= 10 && vc < 30:
		vv = vc + ( vc * 0.5 )
		fmt.Printf( " Valor de venda será de R$ %f ", vv)

	case vc >= 30 && vc < 50:
	    vv = vc + ( vc * 0.4 ) 
		fmt.Printf( " O valor de venda será de R$ %f ", vv)

	case vc >= 50:
		vv = vc + ( vc * 0.3 )
		fmt.Printf( " O valor de venda será de R$ %f ", vv)

	}
}
