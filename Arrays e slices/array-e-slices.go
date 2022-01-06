package main

import "fmt"

func main() {
	//forma de declarar
	var array1 [5]string
	array1[0] = "oi"
	// forma de declarar inplicita
	array2 := [2]string{"Gabriel", "Debora"}
	//outra forma
	array3 := [...]string{"Gabriel", "Debora", "paulino"}
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	//um Array é homogenio é de tamanho fixo;

	//explicitamente forma de Declarar
	var slice1 []string
	slice1 = append(slice1, "Debora Linda")
	//implicitamente forma de declara
	slice2 := []string{
		"gabriel", "Due",
	}
	fmt.Println(slice1)
	fmt.Println(slice2)
	// é como se fosse um ponteiro pra um array veja um exemplo
	slice2 = array3[1:3]
	fmt.Println(slice2)
	array3[2] = "Gabriel Alterado"
	fmt.Println(slice2)
	//Perceba que al mudar o Array3  o slice2 também mudou pois ele estava apontando para o array3
	// em outras palavras slice vem de fatia ou seja slice é uma fatia de um array

	//-------------------------------------//
	//Arrays Internos
	fmt.Println("------Array Internos----------")
	// parametros da func make: 1 tipo, 2 tamanho, 3 capacidade
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade
	//perceba que o slice cria um slice com 10 posicções e capacidade para 11
	// mas oque acontece se estourarmos essa capacidade ? um slice n tem tamanho indefinido ?
	slice3 = append(slice3, 43, 43)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade
	//perceba que ao estourar a capacidade maxima do slice ele criar um novo array com o dobro de posições e aponta para ele

}
