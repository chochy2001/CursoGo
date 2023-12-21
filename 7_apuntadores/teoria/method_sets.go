package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type square struct {
	side float64
}

type shape interface {
	area() float64
	perimeter() float64
}

// Funcionarian con tipos de datos que son apuntadores
// Puede recibir valores de tipo *circle
func (c *circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (c *circle) perimeter() float64 {
	return math.Pi * c.radius * 2
}

// Funcionarian con ambos tipos de datos
// sin apuntador y con apuntador
// Puede recibir valores de tipo square y *square
func (c square) area() float64 {
	return math.Pow(c.side, 2)
}
func (c square) perimeter() float64 {
	return c.side * 4
}

func info(s shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}

func main() {
	c := circle{radius: 5}
	info(&c)
	w := circle{radius: 5}
	//Se cambio para hacer que funcione el codigo
	info(&w)
	/*
		s := square{side: 5}
		info(s)
		x := square{side: 6}
		info(&x)

	*/
}
