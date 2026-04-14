package main
//BUSCA BINÁRIA
import "fmt"

func main() {
	var x int
	fmt.Println(" Qual número você quer encontrar?: ")
	fmt.Scan(&x)
	numeros := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	indice := buscabinaria(numeros, x)

	if indice != -1 {
		fmt.Printf(" O número desejado está no índice %d da lista.\n ", indice)
	} else {
		fmt.Println(" O número solicitado não existe. ")
	}

}

func buscabinaria(numeros []int, chave int) int {
	e := 0
	d := len(numeros) - 1

	for e <= d {
		meio := (e + d) / 2

		if numeros[meio] == chave {
			return meio
		} else if numeros[meio] < chave {
			e = meio + 1
		} else {
			d = meio - 1
		}
	}
	return -1

}
