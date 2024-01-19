// Objetivo: Crear una función genérica Minimo en Go que
// encuentre el valor mínimo en un slice de cualquier tipo numérico.

// Descripción: Este ejercicio requiere que se implemente
// una función genérica Minimo que tome un slice de cualquier tipo numérico
// (como int, float64, etc.) y devuelva el valor mínimo de ese slice.
// La función debe ser aplicable a cualquier tipo de datos numéricos y
// manejar slices vacíos adecuadamente.

// Requisitos:
// La función debe ser capaz de manejar diferentes tipos numéricos.
// Utilizar una restricción de tipo genérico que permita solo tipos numéricos.
// La función debe retornar un valor adecuado (como un valor por defecto
// o un error) si el slice está vacío.
package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Nums interface {
	constraints.Integer | constraints.Float
}

func Minimo[T Nums](slice []T) (T, error) {
	if len(slice) == 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("El slice esta vacio")
	}

	minValue := slice[0]
	for _, v := range slice[1:] {
		if v < minValue {
			minValue = v
		}
	}
	return minValue, nil

}

func main() {

	sliceInts := []int{1, 2, 3, 4, 5}
	minInt, err := Minimo(sliceInts)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Minimo de int: ", minInt)
	}

	var sliceVacio []float32
	minFloat, err := Minimo(sliceVacio)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Minimo de float32: ", minFloat)
	}

}
