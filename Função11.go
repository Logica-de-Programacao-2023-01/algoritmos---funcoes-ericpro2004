package main

import (
	"errors"
	"fmt"
	"math"
)

// Função que recebe um número inteiro como parâmetro e retorna verdadeiro se ele for um número primo e falso caso contrário.
func isPrimo(num int) (bool, error) {
	if num < 0 {
		return false, errors.New("Número negativo não é válido")
	}

	if num < 2 {
		return false, nil
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

func main() {
	numero := 17
	primo, err := isPrimo(numero)
	if err != nil {
		fmt.Println(err)
		return
	}

	if primo {
		fmt.Println(numero, "é um número primo")
	} else {
		fmt.Println(numero, "não é um número primo")
	}
}
