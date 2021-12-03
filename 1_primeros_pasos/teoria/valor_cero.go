package main

import "fmt"

var nombre2 string
var edad2 uint8
var peso float32
var tarea bool
var complejo complex128

func main() {
	fmt.Println(nombre2)
	fmt.Printf("%T, %v,%#v,%%,\n", nombre2,nombre2,nombre2,nombre2)
	fmt.Println(edad2)
	fmt.Printf("%T\n", edad2)
	fmt.Println(peso)
	fmt.Printf("%T\n", peso)
	fmt.Println(tarea)
	fmt.Printf("%T\n", tarea)
	fmt.Println(complejo)
	fmt.Printf("%T\n", complejo)

}
