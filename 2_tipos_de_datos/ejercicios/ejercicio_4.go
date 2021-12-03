// * Escribe un programa que asigne valores a una variable de tipo entero, que el valor lo ingrese el usuario
// * Imprime sus valores en pantalla (decimal, binario, octal y hexadecimal)
// * Usa el bitShifting de 1 a la izquierda y ese valor asignaselo a una nueva variable
// * Imprime los valores en pantalla de la nueva variable en (decimal, binario, octal y hexadecimal)

// ? Explica que es el bitShifting y para que se usa

package main

import "fmt"

func main(){

	var numero int
	fmt.Println("Ingrese un numero: ")
	fmt.Scanf("%d", &numero)
	fmt.Printf("Tu número es: %d,%b,%o,%x\n", numero, numero, numero, numero)
	numeroIzquierda := numero<<1
	fmt.Printf("El número con el bitShifting a la izquierda: %d,%b,%o,%x\n", numeroIzquierda, numeroIzquierda, numeroIzquierda, numeroIzquierda)
}
