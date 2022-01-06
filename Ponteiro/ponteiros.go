package main

import "fmt"

func main() {
	num1 := 3
	num2 := num1
	fmt.Println(num1, num2)
	num1++
	//perceba que ao incremente o num1o valor de num2 nao muda pois num2 recebeu uma copia de num1 e nao uma referencia de memoria
	fmt.Println(num1, num2)

	number1 := 4
	var number2 *int = &number1
	//perceba que o number2 não recebeu o valor de number1 mas sim o endereço da memoria
	fmt.Println(number1, number2)
	// para pegar o valor na memoria colocas o * na frente
	fmt.Println(number1, *number2)

	number1++
	//agr sempre que nnumber 1 mudar number2 muda junto pois ele nao é uma copia com sim uma referencia de memoria
	fmt.Println(number1, *number2)
}
