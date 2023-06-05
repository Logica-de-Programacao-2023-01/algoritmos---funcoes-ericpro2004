package main

import (
	"fmt"
	"strings"
)

// Função que recebe um slice de strings e retorna a concatenação de todas as strings.
func concatenarStrings(strs []string) string {
	return strings.Join(strs, "")
}

func main() {
	slice := []string{"Olá", ", ", "mundo!"}
	concatenacao := concatenarStrings(slice)
	fmt.Println(concatenacao)
}
