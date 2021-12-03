package main

import "fmt"

func main() {
	//Imprime los años que has estado vivo desde tu nacimiento hasta el añoNacimiento actual
	//Usando
	//bucle for candicion{}
	//bucle for{}
	añoNacimiento := 2000
	añoActual := 2050
	contador := 0

	/*
		for añoNacimiento <= añoActual {
			fmt.Println(añoNacimiento)
			añoNacimiento++
			contador++
		}
		fmt.Println("Has estado vivo", contador, "años")
		fmt.Printf("Has estado vivo %d años", contador)

	*/

	for {
		if añoNacimiento > añoActual {
			break
		}
		fmt.Println(añoNacimiento)
		añoNacimiento++
		contador++
	}
	fmt.Println("Has estado vivo", contador, "años")

}
