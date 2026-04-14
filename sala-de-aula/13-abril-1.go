package main
//BUSCA SEQUENCIAL
import "fmt"

func main() {
	var x int
	fmt.Println(" Matricula: ")
	fmt.Scan(&x)
	matricula := []int{10, 20, 30, 40}
	indice := busca(matricula, x)

	if indice != -1 {
		fmt.Printf(" O aluno foi aprovado, seu índice é: %d \n", indice)
	} else {
		fmt.Println(" O aluno não foi encotrado")
	}

}

func busca(matricula []int, chave int) int {
	for i, valor := range matricula {
		if valor == chave {
			return i
		}
	}
	return -1

}
