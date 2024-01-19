//Objetivo: Implementar una función genérica AgruparPorClave en Go
//que agrupe elementos de un slice según una clave proporcionada por una función.

//Descripción: En este ejercicio, deberas crear una
//función genérica que tome un slice de cualquier tipo y una función
//que determine una clave para cada elemento del slice. La función
//AgruparPorClave deberá devolver un mapa cuyas claves son los valores
//retornados por la función de clave y cuyos valores son slices de los
//elementos agrupados.

//Requisitos:
//La función debe ser genérica y aplicable a cualquier tipo de slice.
//La función de clave también debe ser genérica.
//El resultado debe ser un mapa que agrupe los elementos según
//las claves determinadas por la función de clave.

package main

import "fmt"

func AgruparPorClave[T any, K comparable](slice []T, claveFunc func(T) K) map[K][]T {
	resultado := make(map[K][]T)
	for _, v := range slice {
		clave := claveFunc(v)
		resultado[clave] = append(resultado[clave], v)
	}
	return resultado
}

type Persona struct {
	Nombre string
	Edad   int
}

func main() {
	personas := []Persona{
		{"Jorge", 23},
		{"Maria", 25},
		{"Juan", 23},
		{"Pedro", 25},
		{"Ana", 23},
		{"Luis", 25},
		{"Miguel", 50},
	}

	agrupados := AgruparPorClave(personas, func(p Persona) int {
		return p.Edad
	})

	for edad, grupo := range agrupados {
		fmt.Printf("Edad %d: %v\n", edad, grupo)
	}

}
