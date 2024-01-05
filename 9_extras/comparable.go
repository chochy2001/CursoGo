package main

import "fmt"

//Comparable types
//Numeric types
//String
//Boolean
//Pointer types

//non-comparable types
//Slice
//Map
//Function

type Point struct {
	X, Y int
}

func main() {

	point := make(map[Point]string)
	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 3, Y: 4}

	point[p1] = "Hola"
	point[p2] = "mundo"

	p, exists := point[Point{1, 2}]
	fmt.Println("Punto {1,2} existe: ", p, exists)
	a := 5
	b := 6
	if a == b {
		fmt.Println("Son iguales")
	} else {
		fmt.Println("Son diferentes")
	}

}
