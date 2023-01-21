package main

import "fmt"

//Crea una función que retorne una función,
//la función retornada asignasela a una variable y ejecuta esa función

func main() {
	pruebaRetornoFuncion := prueba()
	pruebaRetornoFuncion()
	fmt.Printf("%T\n", prueba)
	fmt.Printf("%T\n", prueba())
	fmt.Printf("%T\n", pruebaRetornoFuncion)
	fmt.Printf("%T\n", pruebaRetornoFuncion())
	fmt.Printf("%v\n", pruebaRetornoFuncion())

}

func prueba() func() int {
	return func() int {
		return 21
	}
}
