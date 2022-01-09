package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		//perceba que as menssagem estão sendo imprimidas igualmente na tela
		// mesmo o canal 1 manda uma menssagem a cada 0,5s enquando a do canal dois
		// manda uma manssagem a cada 2s por causa do bloqueio do canal 2 para mandar a manssagem
		//mensagem1 := <-canal1
		//fmt.Println(mensagem1)
		//mensagem2 := <-canal2
		//fmt.Println(mensagem2)

		//perceba que o select tem uma estrutura pareceda com a do switch
		// ele verifica se tem mensagem para ser recebida assim não obrigando a menssagem1 a esperar a menssagem 2
		select {
		case mensagem1 := <-canal1:
			fmt.Println(mensagem1)
		case mensagem2 := <-canal2:
			fmt.Println(mensagem2)
		}

	}
}
