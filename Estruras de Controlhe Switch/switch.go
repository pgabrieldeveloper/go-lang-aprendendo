package main

import "fmt"

func diasDaSemana(dia int) string {
	switch dia {
	case 1:
		return "dom"
	case 2:
		return "seg"
	case 3:
		return "ter"
	case 4:
		return "quar"
	case 5:
		return "quin"
	case 6:
		return "sex"
	case 7:
		return "sab"
	default:
		return "numero Invalido"
	}

}
func diasDaSemana2(dia int) string {
	switch {
	case dia == 1:
		return "dom"
	case dia == 2:
		return "seg"
	case dia == 3:
		return "ter"
	case dia == 4:
		return "quar"
	case dia == 5:
		return "quin"
	case dia == 6:
		return "sex"
	case dia == 7:
		return "sab"
	default:
		return "numero invalido"
	}

}

func main() {

	dia1 := diasDaSemana(5)
	dia2 := diasDaSemana2(10)
	fmt.Println(dia1)
	fmt.Println(dia2)

}
