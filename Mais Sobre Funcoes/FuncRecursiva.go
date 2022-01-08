package main

import "fmt"

//Função recursiva é uma função que chama ela mesma
// nao é muito comum vemos em uma aplicação mas é bom termos o conhecimento caso seja necessario

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println(fibonacci(4))
}
