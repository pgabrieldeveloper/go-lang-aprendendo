package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}
type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	// inferencia explicito
	var u usuario
	u.nome = "gabriel"
	u.idade = 22
	fmt.Println(u)
	// inferencia implicita

	usuario2 := usuario{idade: 21, nome: "Debora"}
	fmt.Println(usuario2)
	endereco := endereco{
		logradouro: "Rua dos bobos",
		numero:     0,
	}
	usuario3 := usuario{"Paulino", 21, endereco}
	fmt.Println(usuario3)

}
