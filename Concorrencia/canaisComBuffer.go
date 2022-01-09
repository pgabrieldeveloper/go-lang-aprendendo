package main

import "fmt"

func main() {
	//canal := make(chan string)
	//
	//canal <- "ola mundo"
	//
	//mensagem := <-canal
	//Perceba que utilizando da forma acima o programa da deadlokc pois ele bloquaa no recebimento da menssagem e não chega
	//a executar o resto do programa nunca

	//canal com buffer
	// percaba abaixo que o canal segue uma fila de entrada e saida é nao trava no recebimento da menssagem
	canal := make(chan string, 2)
	canal <- "ola mundo"
	canal <- "ola mundo 2"
	mensagem := <-canal
	mensagem2 := <-canal
	canal <- "ola mundo 3"
	mensagem3 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)

}
