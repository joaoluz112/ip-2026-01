package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64

	fmt.Print("Digite o valor de x (em radianos): ")
	fmt.Scan(&x)

	cossenoSerie := 0.0
	sinal := 1.0
	fatorial := 1.0
	potencia := 1.0

	for i := 0; i < 20; i++ {
		if i > 0 {
      
			potencia *= x * x

			fatorial *= float64((2*i - 1) * (2 * i))
		}

		cossenoSerie += sinal * (potencia / fatorial)

		sinal *= -1
	}

	cossenoReal := math.Cos(x)

	diferenca := cossenoSerie - cossenoReal

	fmt.Printf("Cosseno pela série: %.10f\n", cossenoSerie)
	fmt.Printf("Cosseno (math.Cos): %.10f\n", cossenoReal)
	fmt.Printf("Diferença: %.10f\n", diferenca)
}
