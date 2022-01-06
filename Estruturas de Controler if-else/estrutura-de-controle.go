package main

import "fmt"

func main() {
	numero := 2
	// perceba que no go nao é obrigadorio por parenteses no if
	if numero > 5 {
		fmt.Println("Maior que 5")
	} else {
		fmt.Println("Menor que 5")
	}
	// if init
	// perceba que a variavel outroNumero foi criar dentro do if limitante assim sue escopo
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("maior que zero")
	} else {
		fmt.Println("meno que zero")
	}
}
