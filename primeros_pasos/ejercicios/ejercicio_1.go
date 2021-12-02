// ? 1)
// * Crea 4 variables de diferentes tipos,
// * asignale a una tu edad,
// * a otra tu nombre,
// * a otra si te esta gustando el curso
// * y por ultimo el numero pi con 4 decimales

// ? 2)
// ! Imprime en pantalla el valor de las variables,
// ! con su tipo,
// ! y su representación en binario, hexadecimal y octal

// * (Recomendaciones o tips)
// Usa el operador de declaración corta
// Usa el tipo de dato adecuado para cada variable
// Usa un formato especifico en las impresiones

package main

import "fmt"

func main(){
	edad := 20
	nombre := "Jorge"
	curso := true
	pi := 3.1416
	//fmt.Println(edad, nombre, curso, pi)
	fmt.Printf("%v, %T, %b, %o, %x\n", edad, edad, edad, edad, edad)
	fmt.Printf("%v, %T, %b, %o, %x\n", nombre, nombre, nombre, nombre,nombre)
	fmt.Printf("%v, %T, %b, %o, %x\n", curso, curso, curso, curso, curso)
	fmt.Printf("%v, %T, %b, %o, %x\n", pi, pi, pi, pi, pi)







}






