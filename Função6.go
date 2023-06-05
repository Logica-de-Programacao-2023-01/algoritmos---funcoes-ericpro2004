package main

import (
	"errors"
	"fmt"
	"strings"
)

// Função que recebe um slice de strings como parâmetro e retorna uma string com todas as strings concatenadas e separadas por vírgulas.
func concatenarS2(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice está vazio")
	}

	resultado := strings.Join(slice, ",")
	return resultado, nil
}

func main() {
	slice := []string{"Olá", "Mundo", "Go"}
	concatenado, err := concatenarS2(slice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(concatenado)
}
