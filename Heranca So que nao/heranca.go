package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
}

type estudante struct {
	pessoa
	curso   string
	periodo uint8
}

func main() {
	p1 := pessoa{"Gabriel", 22}
	e1 := estudante{p1, "Sistemas de informação", 5}

	fmt.Println(p1)
	fmt.Println(e1)
	fmt.Println(e1.idade)

}
