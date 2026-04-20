package main

import (
	"fmt"
)

func main() {
	var id int
	var peso float64

	var idGordo, idMagro int
	var pesoGordo, pesoMagro float64

	for i := 1; i <= 90; i++ {
		fmt.Printf("\nBoi %d\n", i)

		fmt.Print("Número de identificação: ")
		fmt.Scan(&id)

		fmt.Print("Peso: ")
		fmt.Scan(&peso)

		if i == 1 {
			pesoGordo = peso
			pesoMagro = peso
			idGordo = id
			idMagro = id
		} else {
			if peso > pesoGordo {
				pesoGordo = peso
				idGordo = id
			}

			if peso < pesoMagro {
				pesoMagro = peso
				idMagro = id
			}
		}
	}

	fmt.Println("\n--- RESULTADO ---")
	fmt.Printf("Boi mais gordo -> ID: %d | Peso: %.2f\n", idGordo, pesoGordo)
	fmt.Printf("Boi mais magro -> ID: %d | Peso: %.2f\n", idMagro, pesoMagro)
}
