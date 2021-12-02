package main

import "fmt"

func main(){
	x := 10
	y := 20
	var w int
	type NewInt int
	var z NewInt
	z = 10
	w = 10
	hola := 3.1416
	adios := 4

	fmt.Println("x:",x,"y:",y)
	fmt.Println("x == y: ",x == y) // * Igual
	fmt.Println("x != y:",x != y) // * Diferente
	fmt.Println("x < y:",x < y) // * Menor
	fmt.Println("x > y:",x > y) // * Mayor
	fmt.Println("x <= y: ",x <= y) // * Menor o igual
	fmt.Println("x >= y: ",x >= y) // * Mayor o igual

	fmt.Println("w == z",w == z) // * Igual
	fmt.Println(hola == adios)







}
