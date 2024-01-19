//Objetivo: Implementar una función genérica FiltrarYTransformar en
//Go que filtre y transforme elementos de un slice.

//Descripción: debes crear una función genérica
//FiltrarYTransformar que reciba un slice de cualquier tipo y dos
//funciones: una función de filtro y una función de transformación.
//La función de filtro determinará qué elementos deben mantenerse,
//y la función de transformación modificará esos elementos.
//La función principal debe devolver un nuevo slice con los elementos
//filtrados y transformados.

//Requisitos:
//Uso de generics para el tipo de elementos del slice.
//Implementación de una función de filtro y una función de
//transformación, ambas genéricas.
//Manejo correcto de slices para crear el slice resultante.

package main

import "fmt"

func FiltrarYTransformar[T any, U any](slice []T, filtroFunc func(T) bool, transformarFunc func(T) U) []U {
	var resultado []U
	for _, v := range slice {
		if filtroFunc(v) {
			resultado = append(resultado, transformarFunc(v))
		}
	}
	return resultado
}

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6}
	resultado := FiltrarYTransformar(numeros,
		func(n int) bool { //Filtro
			return n%5 == 0
		},
		func(n int) string { //Transformacion
			return fmt.Sprintf("Numero: %d,", n)

		})
	fmt.Println("Resultado: ", resultado)

}
