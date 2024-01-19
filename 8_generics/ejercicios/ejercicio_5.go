//Objetivo: Crear una función genérica ContarElementos en Go que cuente
//las ocurrencias de cada elemento en un slice.

//Descripción: En este ejercicio, debes implementar una
//función genérica ContarElementos que tome un slice de cualquier tipo
//y devuelva un mapa. Este mapa debe tener como claves los elementos del
//slice y como valores la cantidad de veces que cada elemento aparece en el slice.

//Requisitos:
//La función debe ser aplicable a cualquier tipo de datos en el slice.
//El tipo de los elementos del slice debe ser comparable
//(para que puedan ser claves en el mapa).
//El resultado debe ser un mapa que contenga las ocurrencias de cada elemento.

package main

import "fmt"

func ContarElementos[T comparable](slice []T) map[T]int {
	contador := make(map[T]int)
	for _, elemento := range slice {
		contador[elemento]++
	}
	return contador
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	resultado := ContarElementos(slice)

	fmt.Println("Conteo de elementos: ", resultado)

	sliceString := []string{"Hola", "Mundo", "Hola,", "Mundo", "Hola", "Hola", "Hola"}
	resultadoString := ContarElementos(sliceString)
	fmt.Println("Conteo de elementos: ", resultadoString)

}
