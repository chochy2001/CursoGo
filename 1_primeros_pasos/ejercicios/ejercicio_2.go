// * Declara 3 variables, una de tipo entero, otra de tipo cadena y otra de tipo booleano.
// * Con los identificadores de edad, nombre y si ya dejaste tu reseña de 5 estrellas


// ! Puntos a tomar en cuenta
// Las variables tienen que ser variables globales *
// No puedes usar el operador de declaracion corta *
// No puedes declarar e iniciarlizar las variables en la misma linea *
// Imprime el valor asignado a cada variable sin asignar ningun valor,*  asigna los valores * y vuelve a imprimirlos
// como se llama el valor que se le asigna a las variables al ser declaradas (sin ser inicializadas) *

package main

import "fmt"

/*var edad int
var nombre string
var dejasteReseña bool
 */

func main(){
	fmt.Printf("%v, %v, %v\n", edad,nombre,dejasteReseña)
	edad = 20
	nombre = "Jorge"
	dejasteReseña = true
	fmt.Printf("%v, %v, %v\n", edad,nombre,dejasteReseña)
}



