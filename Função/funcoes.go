package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func main() {
	soma := somar(20, 30)
	println(soma)
	var f = func(texto string) {
		fmt.Println(texto)
	}
	f("Gabriel")
	var calculosMatematicos = func(n1, n2 int8) (int8, int8) {
		soma := n1 + n2
		sub := n1 - n2
		return soma, sub
	}
	soma, sub := calculosMatematicos(10, 5)

	fmt.Println(soma, sub)
}
