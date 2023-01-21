package main

import "fmt"

//Crea una función que tenga variables que solo se pueden acceder en esa función.

func main() {
	pruebaVar := funcion("Jorge")
	fmt.Println(pruebaVar())
	apellido = "Perez"
	fmt.Println(pruebaVar())
}

func funcion(name string) func() string {
	apellido := "Salgado"
	return func() string {
		return name + " :)" + apellido
	}
}
