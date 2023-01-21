package main

import "fmt"

/*
Crea una funcion con un identificador parametrosVariables
que reciba parametrosVariables de enteros (...int)
pasale un slice de enteros y retorna la suma de
todos los valores pasados a la funcion

Crea una funcion con identificador variable que reciba []int
y retorna la suma de todos los valores pasados a la funcion
*/

func main() {
	sliceEnteros := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(parametrosVariables(sliceEnteros...))
	fmt.Println(variable(sliceEnteros))
}

func parametrosVariables(si ...int) int {
	suma := 0
	for _, v := range si {
		suma += v
	}
	return suma
}

func variable(si []int) int {
	suma := 0
	for _, v := range si {
		suma += v
	}
	return suma
}
