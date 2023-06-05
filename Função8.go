package main

import (
	"errors"
	"fmt"
)

// Função que recebe um slice de inteiros como parâmetro e retorna um novo slice com apenas os números pares contidos no slice.
func numerosPares(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	pares := make([]int, 0)
	for _, valor := range slice {
		if valor%2 == 0 {
			pares = append(pares, valor)
		}
	}

	return pares, nil
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	resultado, err := numerosPares(slice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resultado)
}
