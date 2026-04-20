package main

import (
	"fmt"
)

func inverter(arr []int, inicio, fim int) {
	if inicio >= fim {
		return
	}

	arr[inicio], arr[fim] = arr[fim], arr[inicio]

	inverter(arr, inicio+1, fim-1)
}

func main() {
	array := []int{1, 2, 3, 4, 5}

	fmt.Println("Original:", array)

	inverter(array, 0, len(array)-1)

	fmt.Println("Invertido:", array)
}
