package main

import (
	"fmt"
	"sort"
)

// Crie uma função que receba um slice de inteiros e retorne o segundo maior valor.
func main() {
	var num int
	slice := []int{}
	for num != -1 {
		fmt.Print("Digite um numero:")
		fmt.Scan(&num)
		slice = append(slice, num)
	}
	x := len(slice)
	sort.Ints(slice)
	resp := numero(slice, int(x))
	fmt.Print("o segundo maior numero é:", resp)
}
func numero(slice []int, x int) int {
	y := x - 2
	return slice[y]
}
