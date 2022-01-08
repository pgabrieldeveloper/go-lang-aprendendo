package main

import (
	"fmt"
	"math"
)

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

type retangulo struct {
	altura  float64
	largura float64
}

func (c retangulo) area() float64 {
	return c.largura * c.altura
}

//as interfaces so aceita assinatura de metodos
type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A area da figura é: %0.2f \n", f.area())
}

func main() {

	circulo := circulo{5}
	retangulo := retangulo{5, 10}

	escreverArea(circulo)
	escreverArea(retangulo)

}
