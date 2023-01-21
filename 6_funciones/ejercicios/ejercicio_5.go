package main

import (
	"fmt"
	"math"
)

//Crea un tipo cuadrado, otro círculo y triángulo
//agrégale un método para calcular el área y que la retorne.

//Crea un tipo figura que defina una interface de
//todo lo que tenga el metodo area.

//Crea una funcion con el identificador info
//que tome cualquier figura e imprima el area.

//Crea un valor para el círculo, tríangulo y cuadrado,
//usa la función info para imprimir sus áreas

type circulo struct {
	radio int
}
type triangulo struct {
	base, altura int
}
type cuadrado struct {
	lado int
}

func (c circulo) area() int {
	radio2 := math.Pow(float64(c.radio), 2)
	area := math.Pi * radio2
	return int(area)
}
func (t triangulo) area() int {
	area := (t.base * t.altura) / 2
	return area
}
func (c cuadrado) area() int {
	//area := c.lado*c.lado
	area := math.Pow(float64(c.lado), 2)
	return int(area)
}

type figura interface {
	area() int
}

func info(f ...figura) {
	//fmt.Printf("%T\n", f)
	for _, v := range f {
		fmt.Printf("Área: %v\n", v.area())
	}
	//fmt.Printf("Área: %v\n", f.area())
}

func main() {
	nCirculo := circulo{
		radio: 2,
	}
	nCuadrado := cuadrado{
		lado: 3,
	}
	nTriangulo := triangulo{
		base:   3,
		altura: 10,
	}
	info(nCirculo, nCuadrado, nTriangulo)
}
