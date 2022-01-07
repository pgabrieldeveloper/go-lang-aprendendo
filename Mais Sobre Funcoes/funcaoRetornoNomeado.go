package main

import "fmt"

//perceba que ao nomear o retorno da função em sua assinatura
// nao precisamos colocar nada no pois o programa ja entende quem sera retornado e nem criar variaveis "soma" e "sub" pois elasa ja foram criadas na assinatura da funcção
func calculosMatematicos(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	soma, sub := calculosMatematicos(20, 10)
	fmt.Println(soma, sub)
}
