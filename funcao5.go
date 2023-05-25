package main

import "fmt"

// Crie uma função que receba um slice de inteiros e um valor inteiro,
// e retorne a posição do primeiro elemento igual ao valor no slice.
// Caso não encontre, retorne -1.
func main() {
	var num int
	x := -1
	slice := []int{}
	for x != 0 {
		fmt.Print("Digite um numero:")
		fmt.Scan(&x)
		slice = append(slice, x)
	}
	fmt.Print("digite o numero e ira retornar a posiçao:")
	fmt.Scan(&num)
	resp := numero(slice, num)
	fmt.Print("Posição:", resp)

}
func numero(slice []int, num int) int {
	for i, n := range slice {
		if n == num {
			return i
		}
	}
	return -1
}
