package main

import "fmt"

//Parámetros de Tipo: Se definen en los corchetes [] después del nombre de la función.
//Estos parámetros actúan como marcadores de posición para los tipos que se pasarán a
//la función.

//Restricciones de Tipo: Las restricciones (any, constraints.Ordered, etc.) especifican
//qué tipos de datos pueden reemplazar los parámetros de tipo. any es la restricción más
//genérica, permitiendo cualquier tipo.

//Uso de Parámetros de Tipo en Funciones: Los parámetros de tipo se utilizan dentro del
//cuerpo de la función como si fueran tipos reales. Go determinará los tipos reales en
//tiempo de compilación.

func AplicarOperacion[T any, U any](slice []T, operacion func(T) U) []U {
	resultado := make([]U, 0, len(slice))
	for _, v := range slice {
		resultado = append(resultado, operacion(v))
	}
	return resultado

}

func main() {
	cuadrados := AplicarOperacion([]int{1, 2, 3}, func(n int) int {
		return n * n
	})
	for _, v := range cuadrados {
		fmt.Println(v)
	}

	cadenas := []string{"manzana", "platano", "pera"}
	longitudes := AplicarOperacion(cadenas, func(s string) int {
		return len(s)
	})
	fmt.Println(longitudes)

}
