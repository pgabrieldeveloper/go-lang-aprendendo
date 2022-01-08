package main

import "fmt"

func inverterNumero(numero int) int {
	return numero * -1
}
func inverterNumeroComPonteiro(numero *int) {
	*numero = *numero * -1
}
func main() {

	numero := 10
	numeroInvertido := inverterNumero(numero)
	//Perceba que ao chamar a função "inverterNumero" a variavel numero nao foi alterada pq oque foi passado para a função foi uma copia da mesma
	fmt.Println(numero, numeroInvertido)
	//Perceba que ao chamar a função "inverterNumeroComPonteiro" a variavel numero tbm é mudada pois
	// oque é passa não é uma copia do numero mas sim sua refecia (seu endereço de memoria)
	inverterNumeroComPonteiro(&numero)
	fmt.Println(numero)

}
