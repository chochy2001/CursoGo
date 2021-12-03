// * Del ejercicio anterior asigna los valores de 50, "Capdesis", true (al tipo de dato que corresponda)

// ! Usa el String print (recuerda del paquete fmt), imprime todos los valores con un solo print *
// ! y en la misma linea (agregale un formato).
// ! Asignaselos a una nueva variable que se llame resultado.
package main

import "fmt"

//var edad int
var nombre string
var dejasteReseña bool

func main(){
	edad = 49
	nombre = "Capdesis"
	dejasteReseña = true

	resultado := fmt.Sprintf("\t%v,\t%v,\t%v\n",edad,nombre,dejasteReseña)
	fmt.Println(resultado)
}

