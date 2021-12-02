// * Escribe un programa que imprima un numero en entero, decimal, binario y hexadecimal.
// ! Si puedes pideselos al usuario.

package main

import "fmt"

func main(){

	var edad uint8
	var nombre string

	fmt.Println("Introduce tu edad: ")
	fmt.Scanf("%d", &edad)
	fmt.Println("Ingresa tu nombre")
	fmt.Scanf("%s",&nombre)
	//fmt.Printf("Hola: %s,%v,%d,%b,%x",nombre, edad, edad, edad, edad)
	fmt.Printf("Hola: %s tu edad es de: %d años",nombre, edad)
}
