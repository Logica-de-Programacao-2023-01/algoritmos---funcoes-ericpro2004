package main

import "fmt"

// Função que conta quantas vogais existem em uma string.
func contarVogais(str string) int {
	vogais := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	count := 0

	for _, char := range str {
		for _, vogal := range vogais {
			if char == vogal {
				count++
				break
			}
		}
	}

	return count
}

func main() {
	texto := "Ola, mundo!"
	qtdVogais := contarVogais(texto)
	fmt.Println("Quantidade de vogais: ", qtdVogais)
}
