package main

import "fmt"

func recupera() {
	if r := recover(); r != nil {
		println("Programa Recuperado Com sucesso !")
	}
}

// o metodo panic da encerra a execução do programa sendo necessario usar o recover() para tratar o panic
// não é a melhor forma de tratar erro no GO mas é bom conhecer caso necessario
func mediaAluno(n1, n2 float64) bool {
	defer recupera()
	media := (n1 + n2) / 2

	if media > 6 {
		fmt.Println("Foi Aprovado")
		return true
	} else if media < 6 {
		fmt.Println("foi Reprovado")
		return false
	}

	panic("A media é exatamente 6")
}

func main() {
	mediaAluno(4, 7)
	mediaAluno(7, 7)
	mediaAluno(6, 6)
}
