//DIVISÃO

func divisao(n1, n2 float64) (bool, float64) {

	error := false
	resultado := 0.0

	if n2 == 0 {
		error = true
	} else {
		resultado = n1 / n2
	}

	return error, resultado
}

func main() {
	var n1, n2 float64

	fmt.Print(" Digite dois números: ")
	fmt.Scanln(&n1, &n2)
	error, resultado := divisao(n1, n2)

	if !error {
		fmt.Printf(" O resultado é %.2f \n", resultado)
	} else {
		fmt.Println(" A divisão não é possível")
	}

}
