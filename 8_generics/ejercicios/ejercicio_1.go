//Objetivo: Implementar una función genérica en Go que pueda intercambiar
//los valores de dos variables de cualquier tipo.

//Descripción: Los generics en Go permiten escribir funciones y tipos que
//son agnósticos respecto al tipo de datos con el que trabajan. En este ejercicio,
//debes implementar una función genérica
//Intercambiar,
//que tome dos punteros a variables de cualquier tipo y los intercambie.

//Requisitos:
//La función debe ser capaz de intercambiar valores de diferentes tipos
//(por ejemplo, int, float64, string, etc.).
//Usar generics para que la función sea aplicable a cualquier tipo.

package main

import "fmt"

func Intercambiar[T any](a, b *T) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {

	x, y := 1, 2
	Intercambiar(&x, &y)
	fmt.Println("Intercambio de int: x = ", x, " y = ", y)

	a, b := 1.1, 2.2
	Intercambiar(&a, &b)
	fmt.Println("Intercambio de float64: a = ", a, " b = ", b)

	c, d := "Hola", "Mundo"
	Intercambiar(&c, &d)
	fmt.Println("Intercambio de string: c = ", c, " d = ", d)

}
