package main

import "fmt"

// Função que faz a Média de números inteiros solicitados em um slice.
func calcularMedia(slice []int) float64 {
	if len(slice) == 0 {
		return 0
	}

	sum := 0
	for _, num := range slice {
		sum += num
	}

	return float64(sum) / float64(len(slice))
}

func main() {
	valores := []int{2, 4, 6, 8, 10}
	media := calcularMedia(valores)
	fmt.Println("Média: ", media)
}
