package main

import (
	"fmt"
)

func main() {

	// for semelhante ao while
	i := 0
	for i < 10 {
		i++
		fmt.Println("incrementando i: ", i)
	}

	//for padrão das linguaguens
	for j := 0; j < 10; j++ {
		fmt.Println("incrementando J:", j)
	}

	//For iteravel semelhante ao forEach
	nomes := [3]string{"gabriel", "debora", "due"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

}
