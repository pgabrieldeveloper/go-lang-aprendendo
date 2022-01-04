package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Escrevendo no main")
	auxiliar.Escrever()
	auxiliar.Escrever2()
	erro := checkmail.ValidateFormat("koymatt@gmail.com")
	fmt.Println(erro)

}
