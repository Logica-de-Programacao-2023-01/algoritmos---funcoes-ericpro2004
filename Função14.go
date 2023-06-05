package main

import (
	"errors"
	"fmt"
)

// Função que recebe dois slices de inteiros como parâmetros e retorna um novo slice contendo apenas os números presentes nos dois slices.
func numerosComuns(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("Slice vazio")
	}

	numerosComuns := make([]int, 0)

	for _, num1 := range slice1 {
		for _, num2 := range slice2 {
			if num1 == num2 {
				numerosComuns = append(numerosComuns, num1)
				break
			}
		}
	}

	return numerosComuns, nil
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}
	resultado, err := numerosComuns(slice1, slice2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resultado)
}
