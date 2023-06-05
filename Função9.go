package main

import (
	"errors"
	"fmt"
	"strings"
)

// Função que recebe uma string como parâmetro e retorna um novo slice com todas as palavras contidas na string.
func extrairPalavras(texto string) ([]string, error) {
	if texto == "" {
		return nil, errors.New("A string está vazia")
	}

	palavras := strings.Split(texto, " ")
	return palavras, nil
}

func main() {
	texto := "Olá, como vai você?"
	resultado, err := extrairPalavras(texto)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resultado)
}
