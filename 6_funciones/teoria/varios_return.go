package main

import "fmt"

//Varios return
//Parametros variables

func main() {
	fmt.Println(variosReturn(102, "Jorge"))
	fmt.Println(unReturn())
	x, y := variosReturnDiferenteDato(21, "Jorge")
	fmt.Println(x, y)
}

func variosReturnDiferenteDato(edad int, nombre string) (valor1 bool, valor2 float64) {
	fmt.Println("Tu edad es:", edad, " Tu nombre es:", nombre)
	return true, 10.5
}

func variosReturn(a int, b string) (valor1 int, valor2 string) {
	fmt.Println(a, b)
	return 10, "hola"
}

func unReturn() (a int) {
	return 10
}
