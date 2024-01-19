//Objetivo: Crear una función genérica en Go que pueda sumar los elementos
//de un slice, independientemente del tipo numérico del slice.

//Descripción: Este ejercicio desafia a implementar una función genérica
//SumarSlice que tome un slice de cualquier tipo numérico
//(por ejemplo, int, float64, etc.) y devuelva la suma de sus elementos.
//El ejercicio introduce el concepto de restricciones de tipo más
//específicas en generics.

//Requisitos:
//La función debe ser capaz de sumar elementos de slices de
//diferentes tipos numéricos.
//Utilizar una restricción de tipo genérico que permita solo
//tipos numéricos.
//Considerar la precisión en el caso de tipos de punto flotante.

package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Num interface {
	constraints.Integer | constraints.Float
}

func SumarSlice[T Num](slice []T) T {
	var suma T
	for _, v := range slice {
		suma += v
	}
	return suma
}

func main() {

	sliceInt := []int{1, 2, 3, 4, 5}
	sumaInt := SumarSlice(sliceInt)
	println("Suma de int: ", sumaInt)

	sliceFloat := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	sumaFloat := SumarSlice(sliceFloat)
	fmt.Printf("Suma de float64: %.5f\n", sumaFloat)

	/*
		sliceString := []string{"Hola", "Mundo"}
		sumaString := SumarSlice(sliceString)
		println("Suma de string: ", sumaString)
	*/

}
