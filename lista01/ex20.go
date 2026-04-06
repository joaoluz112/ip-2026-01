package main

import (
	"fmt"
)

func main() {
	var horas, minutos, segundos int
	var total int

	fmt.Scan(&horas)
	fmt.Scan(&minutos)
	fmt.Scan(&segundos)

	total = horas*3600 + minutos*60 + segundos

	fmt.Printf("O TEMPO EM SEGUNDOS E = %d\n", total)
}
