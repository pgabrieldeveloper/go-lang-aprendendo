package main

import "fmt"

func main() {
	//Maps

	//como criar
	map1 := map[string]string{"nome": "gabriel", "sobrenome": "matias"}
	fmt.Println(map1["nome"])

	//Map de map
	map2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Gabriel",
			"segundo":  "matias",
		},
	}
	fmt.Println(map2["nome"])
	fmt.Println(map2["nome"]["segundo"])
	fmt.Println(len(map2))
	//para dedletar uma chave do map
	delete(map2, "nome")
	fmt.Println(len(map2))
	//perceba que o novo tamanho do map agr é 0

	//para adicionar um novo campo ao map e semelhante a um Array

	map2["curso"] = map[string]string{
		"nome":    "Sistemas de Informação",
		"periodo": "8",
	}
	fmt.Println(map2)
	// map  é uma estrutura semelhante a dicionarios

}
