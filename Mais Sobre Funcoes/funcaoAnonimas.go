package main

import "fmt"

func main() {
	//Função anonimas Exemplos perceba que o "()" no final da escrita da func
	// pedi para rodar ela assim que for criada
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebendo Func Anonimo -> %s", texto)
	}("Gabriel")

	fmt.Println(retorno)
}
