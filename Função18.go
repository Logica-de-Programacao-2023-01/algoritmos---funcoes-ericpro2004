package main

import (
	"errors"
	"fmt"
)

// Função que recebe um slice de inteiros e uma função como parâmetros.
func APFUNC(slice []int, fn func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("Slice vazio")
	}

	resultado := 0
	for _, valor := range slice {
		resultado += fn(valor)
	}

	return resultado, nil
}

func dobrar(valor int) int {
	return valor * 2
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	resultado, err := APFUNC(slice, dobrar)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resultado)
}
