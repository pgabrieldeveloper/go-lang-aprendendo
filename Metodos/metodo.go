package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

//Assinatura de uma metodo "func ('Tipo do struct') 'nome do metodo()'"
func (u usuario) salvar() {
	fmt.Printf("Salvando usuario: %s no banco de dados\n", u.nome)
}

//se quiser mudar algum dado do struct passo o tipo como referencia (ponteiro como no exemplo abaixo)
func (u *usuario) fazerAniversario() {
	u.idade++
}

// a principal diferença entre função e metodo e que metodo esta ligado a algo como um struct
func main() {

	usuario1 := usuario{"Gabriel", 18}
	fmt.Println(usuario1)
	usuario1.salvar()
	fmt.Println(usuario1.idade)
	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)

}
