package main

import "fmt"

//O defer adia a execução do programa ate o ultimo momento possivel
//Perceba quqe ao executar a função ele so da o print Antes do return
//assim nao precisando  colocar o print antes de cada Return
func mediaAluno(n1, n2 float64) bool {
	defer println("Media Calculada o Resultado será retornado ! ")
	if (n1+n2)/2 > 6 {
		return true
	}
	return false
}

func main() {

	fmt.Println(mediaAluno(5, 9))

}
