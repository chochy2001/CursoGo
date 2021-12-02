package main

import "fmt"

var edad uint8       // ! 0
var comidaFav string // ! ("")
var feliz bool       // ! false
var flotante float32 // ! 0

func main() {
	//declara una variable y asigna el valor al mismo tiempo

	// ! Operador de declaracion corta
	nombre := "Jorge"

	fmt.Println(nombre, edad, feliz, flotante)
	edad = 20
	feliz = true

	var apellido = "Chochy"

	fmt.Println(nombre, apellido, edad, feliz)
	comidaFav = "Pizza"

	fmt.Println(nombre, apellido, edad, comidaFav)
	imprimirEdad()

}

func imprimirEdad() {
	fmt.Println("Mi edad es:", edad, "Mi comida favorita es:", comidaFav)
}
