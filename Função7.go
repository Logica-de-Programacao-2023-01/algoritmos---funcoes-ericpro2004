package main

import (
	"errors"
	"fmt"
)

// Uma função que recebe um slice de inteiros e uma função como parâmetros.
func aplicarFuncao(slice []int, funcao func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	resultado := make([]int, len(slice))
	for i, valor := range slice {
		resultado[i] = funcao(valor)
	}

	return resultado, nil
}

func dobro(numero int) int {
	return numero * 2
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	resultado, err := aplicarFuncao(slice, dobro)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resultado)
}
