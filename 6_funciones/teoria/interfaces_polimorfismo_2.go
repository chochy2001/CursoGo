package main

import "fmt"

func main() {

	var hola int = 21
	fmt.Println(hola)
	fmt.Printf("%T\n", hola)

	type edad int
	var x edad = 12
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	hola = int(x)
	fmt.Println(hola)
	fmt.Printf("%T\n", hola)

}
