package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multplexador(escrever("Gabriel"), escrever("Debora"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

//o multiplexador sador usa o select para filtrar as menssagems prontas e devolver conforme elas estão saindo
func multplexador(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)
	go func() {
		for {
			select {
			case menssagem := <-canalEntrada1:
				canalDeSaida <- menssagem
			case menssagem := <-canalEntrada2:
				canalDeSaida <- menssagem
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor do texto: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal

}
