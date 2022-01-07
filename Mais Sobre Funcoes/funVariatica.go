package main

import (
	"fmt"
)

func soma(numeros ...int) int {
	total := 0
	//percena que ao botar os "...int" ele pega do 1 ate o ultimo paremetro e devolvendo um slice
	//sendo assim podes iterar sobre ele para develer o valor total da soma
	//O valor variatico so pode ter 1 por função e sempre tem que ser o ultimo parametro
	for _, numero := range numeros {
		total += numero
	}
	return total

}

func main() {
	totalSoma := soma(1, 2, 3, 4)
	fmt.Println("soma total:", totalSoma)
}
