package main

import "time"

func escrever(texto string, canal chan string) {
	defer close(canal)
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		canal <- texto
	}
}
func main() {
	// Sintax para criar um canal
	canal := make(chan string)

	go escrever("gabriel", canal)

	for {
		//perceba que acima foi usado o go é o programa segue sme experar a execução da função
		//acabar porem como a variavel mensagem esta esperando uma mensagem do canal
		//o programa so volta a executar depois que a mensagem seja entregue
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		println(mensagem)
	}

}
