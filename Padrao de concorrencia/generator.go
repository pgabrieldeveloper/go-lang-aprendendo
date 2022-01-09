package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Ocolus")

	for i := 0; i < 5; i++ {
		fmt.Println(<-canal)
	}

}

//o padrão generator encapsula toda a  complexidade  do concorrencia da func main
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
