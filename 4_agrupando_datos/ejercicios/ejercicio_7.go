package main

import "fmt"

func main() {
	//Crea un slice de dos dimensiones de tipo uint16
	//agregale los años de nacimiento de tus amigos
	// que los datos los pueda ingresar el usuario. de manera grafica tiene que verse así
	//[[123,456,789]],
	//[[234,567,890]]
	//Imprime el indice y el valor de los elementos del slice

	slice1 := []uint16{123, 456, 789}
	slice2 := []uint16{234, 567, 890}
	fmt.Println(slice1, slice2)
	sliceDosDimensiones := [][]uint16{slice1, slice2}
	fmt.Println(sliceDosDimensiones)

	/*
		for indice,valor := range sliceDosDimensiones{
			fmt.Println("Indice:",indice,"Valor:",valor)
			for i,v := range valor{
				fmt.Println("Indice:",i,"Valor:",v)
			}
		}
	*/

	for i := 0; i < len(sliceDosDimensiones); i++ {
		fmt.Println("Indice:", i, "Valor:", sliceDosDimensiones[i])
		for j := 0; j < len(sliceDosDimensiones[i]); j++ {
			fmt.Println("Indice:", j, "Valor:", sliceDosDimensiones[i][j])
		}
	}

}
