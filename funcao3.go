package main

import "fmt"

//Crie uma função que receba um slice de strings e retorne a concatenação de todas as strings.

func main() {
	n1 := "ola"
	n2 := ",mundo"
	resp := concatena(string(n1), string(n2))
	fmt.Print(resp)
}
func concatena(n1, n2 string) string {
	return n1 + n2
}
