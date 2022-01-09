package main

import "fmt"

//função fibonacci
func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

//função worker
//a função worker recebe 2 canais como parametro um de tarefaz e um de resultados
//cada tarefa é um numero de para adiciona a fila de calcula de fibonacci
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)
	//coloquei 4 "go work() por limitações da minha maquina de 4 processadores"
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	// ao chegar nessa linhas as tarefas sao adicionadas que serão os numeros adicionados para a fila de calculo de fibonacci
	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)
	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}
