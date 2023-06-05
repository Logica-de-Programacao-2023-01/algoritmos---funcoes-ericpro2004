package main

import (
	"fmt"
	"sort"
)

// Função que recebe um slice de inteiros e retorna o segundo maior valor.
func segundoMaiorValor(nums []int) (int, error) {
	if len(nums) < 2 {
		return 0, fmt.Errorf("O slice deve ter pelo menos dois elementos")
	}

	sort.Ints(nums)
	return nums[len(nums)-2], nil
}

func main() {
	slice := []int{10, 5, 8, 12, 3}
	segundoMaior, err := segundoMaiorValor(slice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(segundoMaior)
	}
}
