package main

func escrever(texto string) {
	for {
		println("Escrevendo 1", texto)
	}

}

func escrever1(texto string) {
	for {
		println("Escrevendo 2", texto)
	}

}

//Perceba que se deixar o programa seguir o fluxo normal a função escrever1 nuncna sera executada
//Pois a função escrever esta em um loop infinito
//Perceba que usando o goroutine (go na frente da função)
// basicamente estamos falando para o programa seguir a execução normal sem esperar que a função termine
func main() {
	go escrever("Gabriel") //goroutine
	escrever1("Debora")
}
