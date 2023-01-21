package main

import "fmt"

//Crea una función que su identificador sea suma y otra que sea resta.
//suma debe regresar un entero
//resta debe regresar un entero y un booleano

func main() {
	fmt.Println(suma())
	fmt.Println(resta())
	varSuma := suma()
	varResta, _ := resta()
	varResta1, varBool := resta()
	_, varBool1 := resta()
	fmt.Println(varSuma)
	fmt.Println(varResta)
	fmt.Println(varResta1, varBool)
	fmt.Println(varResta, varBool1)
}

func suma() int {
	return 1
}

func resta() (int, bool) {
	return 2, false
}
