package main

import (
	"errors"
	"fmt"
)

func main() {
	const nome string = "gabriel"
	fmt.Println(nome)
	const numero int8 = 40
	fmt.Println(numero)
	const numeroReal float32 = 12.4
	fmt.Println(numeroReal)
	char := 'B'
	fmt.Println(char)
	const boleano bool = true
	fmt.Println(boleano)
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
