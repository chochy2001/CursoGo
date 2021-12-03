package main

import "fmt"

var nombre = "Juan"
var edadNombre uint8 = 30

func main() {
	comidaFav := "Pizza \"asdf\" "
	lugarFav := `Mi 
			lugar 
favorito 				es mi				 "cuarto"
`
	fmt.Println(nombre)
	fmt.Printf("%T\n", nombre)
	fmt.Println(edadNombre)
	fmt.Printf("%T\n", edadNombre)
	fmt.Println(comidaFav)
	fmt.Printf("%T\n", comidaFav)
	fmt.Println(lugarFav)
	fmt.Printf("%T\n", lugarFav)
	// ! como es un lenguaje de tipado estatico, no se puede cambiar el tipo de dato
}
