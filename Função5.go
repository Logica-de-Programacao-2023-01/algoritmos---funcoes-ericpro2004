package main

import "fmt"

// Função que recebe um slice de inteiros e um valor inteiro, e retorna a posição do primeiro elemento igual ao valor no slice.
func encontrarPosicao(slice []int, valor int) int {
	for i, num := range slice {
		if num == valor {
			return i
		}
	}
	return -1
}

func main() {
	slice := []int{10, 5, 8, 12, 3}
	valor := 8
	posicao := encontrarPosicao(slice, valor)
	fmt.Println(posicao)
}
