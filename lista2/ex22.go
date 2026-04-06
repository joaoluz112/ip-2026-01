package main

import (
	"fmt"
)

func main() {
	var matricula int
	var horasExtras float64

	const salarioMinimo = 788.0
	const valorHoraExtra = 10.0

	fmt.Print("Digite a matrícula do funcionário: ")
	fmt.Scan(&matricula)

	fmt.Print("Digite a quantidade de horas-extras: ")
	fmt.Scan(&horasExtras)

	
	salarioExtra := horasExtras * valorHoraExtra
	salarioBruto := 3*salarioMinimo + salarioExtra

	var descontoINSS float64
	var descontoIR float64

	
	if salarioBruto > 1500 {
		descontoINSS = salarioBruto * 0.12
	}

	
	if salarioBruto > 2000 {
		descontoIR = salarioBruto * 0.20
	}

	salarioLiquido := salarioBruto - descontoINSS - descontoIR

	
	fmt.Println("\n--- Dados do Funcionário ---")
	fmt.Printf("Matrícula: %d\n", matricula)
	fmt.Printf("Horas-extras: %.2f\n", horasExtras)
	fmt.Printf("Salário bruto: R$ %.2f\n", salarioBruto)
	fmt.Printf("Desconto INSS: R$ %.2f\n", descontoINSS)
	fmt.Printf("Desconto IR: R$ %.2f\n", descontoIR)
	fmt.Printf("Salário líquido: R$ %.2f\n", salarioLiquido)
}
