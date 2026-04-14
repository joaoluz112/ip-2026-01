package main
// FATORIAL RECURSIVO
import "fmt"

func main() {

	fmt.Println(fatorial(7))

}

func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}
