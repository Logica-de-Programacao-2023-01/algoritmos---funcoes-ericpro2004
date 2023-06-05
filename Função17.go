package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// Função que recebe um slice de strings como parâmetro e retorna uma nova string com as strings em ordem alfabética.
func ordenarStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Slice vazio")
	}

	sort.Strings(slice)
	resultado := strings.Join(slice, ", ")

	return resultado, nil
}

func main() {
	slice := []string{"banana", "abacaxi", "laranja", "uva"}
	resultado, err := ordenarStrings(slice)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resultado)
}
