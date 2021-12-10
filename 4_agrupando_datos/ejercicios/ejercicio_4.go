package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//agrega el elemento 11 al final del slice *
	//Imprime el slice
	//agrega los valores del slice a un nuevo slice con valores agregados de 12,13,14,15
	//Imprime el nuevo slice

	fmt.Println(slice)
	slice = append(slice, 11)
	fmt.Println(slice)
	slice2 := append(slice, 12, 13, 14, 15)
	fmt.Println(slice2)
	fmt.Println(slice)

}
