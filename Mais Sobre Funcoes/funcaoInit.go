package main

import "fmt"

var n int

//Função init é carregada antes da função main é pode ser criada uma por arquivo
func init() {
	fmt.Println("Função init")
	n = 10
}

func main() {
	fmt.Println("Funcao main", n)
}
