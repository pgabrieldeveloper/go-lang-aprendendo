package main

import (
	"sync"
	"time"
)

func escrever() {
	for i := 0; i < 4; i++ {
		println("Olá Mundo !")
		time.Sleep(time.Second)

	}

}

func escrever1() {
	for i := 0; i < 4; i++ {
		println("Programando em GO !")
		time.Sleep(time.Second)
	}

}

func main() {
	//Declarando uma variavel waitGroup
	var waitGroup sync.WaitGroup
	//Adicionando quantidades de Goroutine
	waitGroup.Add(2)

	go func() {
		escrever()
		//Retira um da fila de Goroutines
		waitGroup.Done()
	}()

	go func() {
		escrever1()
		//Retira um da fila de Goroutines
		waitGroup.Done()
	}()
	//Espera todas as A fila esta zerada para proseguir o programa
	waitGroup.Wait()
}
