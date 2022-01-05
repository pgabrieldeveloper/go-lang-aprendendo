package main

import "fmt"

func main() {
	// aritmetricos
	soma := 1 + 1
	subtracao := 2 - 1
	mult := 4 * 2
	div := 4 / 2
	restoDivisao := 4 % 3
	fmt.Println(soma, subtracao, mult, div, restoDivisao)

	// atribuição
	const nome string = "gabriel"
	nome1 := "gabriel"

	// relacionais
	//padrão de qualquer linguagem
	// <, > , == , != , <= ,>=

	//operadores logicos
	//padrão de qualquer linguagem
	// E: &&, OU ||, Negação !

	//operadores unarios
	numero := 4
	numero += 2
	numero -= 2
	numero /= 1
	numero *= 2
	numero++
	numero--

	// operador ternario
	// Em go não existe operador ternario
	//texto := numero > 5 ? "maior que cinco" : "menor que cinco"
}
