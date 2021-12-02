package main

import "fmt"

func main() {
	var x int
	type edad int
	var y edad
	x = 20
	y = 31
	fmt.Printf("Tipo: %T, Valor: %v\n", x,x)
	fmt.Printf("Tipo: %T, Valor: %v", y,y)
	// x = y
	// ! Eso no se puede hacer ya que son de diferente tipo

	// * Esto si se puede hacer
	x = int(y)
	fmt.Println("\nEl nuevo valor de x =",x)
	y = edad(x)
	fmt.Println("\nEl nuevo valor de y =",y)

	//Ejemplo de conversiones
	var entero int = 10
	var flotante float64 = 20.9999
	var cadena string = "Hola"

	fmt.Println(entero,flotante,cadena)
	entero = int(flotante)
	fmt.Println(entero)
	entero = 10
	flotante = float64(entero)
	fmt.Println(flotante)
}
