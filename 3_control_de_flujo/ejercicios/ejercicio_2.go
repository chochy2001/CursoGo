package main

import "fmt"

func main() {
	//Imprimir las letras del abecedario de esta manera
	//97 a 122
	//a-z
	//salida:
	//97
	//a
	//a
	//a
	//a
	//98
	//b
	//b
	//b
	//b
	//99
	//c
	//c
	//c
	//c
	for i := 97; i <= 122; i++ {
		fmt.Println(i)
		for j := 0; j <= 3; j++ {
			fmt.Printf("\t%c\n", i)
			//fmt.Println("\t",string(i))
		}
	}

	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j <= 3; j++ {
			fmt.Printf("\t%c\n", i)
			//fmt.Println("\t",string(i))
		}
	}

}
