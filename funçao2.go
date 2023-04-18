package main

import "fmt"

func main() {
	var texto string
	var cont int
	fmt.Println("conjunto de vogais")
	fmt.Print("digite uma palavra:")
	fmt.Scan(&texto)
	vogal := vogais(string(texto), int(cont))
	fmt.Printf("A um total de %d vogais", vogal)
}
func vogais(texto string, cont int) int {
	for i := 0; i < len(texto); i++ {
		if texto[i] == 'a' || texto[i] == 'e' || texto[i] == 'i' || texto[i] == 'o' || texto[i] == 'u' {
			cont++
		}
	}
	return cont
}
