package main

import (
	"fmt"
)

func main() {
	var lista int
	fmt.Println("media de uma lista")
	fmt.Print("digite o tamanho da lista")
	fmt.Scan(&lista)
	slice := make([]int, lista)
	total := media(slice[lista], lista)
	fmt.Printf("A media dos %d numeros é %d", lista, total)
}
func media(slice int, lista int) int {
	var soma int
	for i := 0; i < lista; i++ {
		fmt.Printf("digite o numero %d:", i)
		fmt.Scan(&slice)
		soma += slice
	}
	return soma / lista
}
