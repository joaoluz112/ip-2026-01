package main

import (
	"fmt"
)

func main() {
	var casos int
	fmt.Scan(&casos)

	for i := 1; i <= casos; i++ {
		var publico int
		var pPop, pGer, pArq, pCad float64

		fmt.Scan(&publico, &pPop, &pGer, &pArq, &pCad)

		// Quantidade de pessoas por categoria
		popular := (pPop / 100) * float64(publico)
		geral := (pGer / 100) * float64(publico)
		arquibancada := (pArq / 100) * float64(publico)
		cadeiras := (pCad / 100) * float64(publico)

		// Cálculo da renda
		renda := popular*1 + geral*5 + arquibancada*10 + cadeiras*20

		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i, renda)
	}
}
