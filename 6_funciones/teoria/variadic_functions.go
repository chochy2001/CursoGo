package main

import "fmt"

//Funciones con parametros variables
//(Variadic Parameters)

func main() {
	//recibirEnteros("Nombre", true, 1)
	//recibirEnteros("NOMBRE", false, 2, 3, 4, 5, 6, 7, 8, 9)
	valorFinal := sumaNumeros1(1, 2, 3, 4, 5, 6)
	fmt.Println("El valorFinal:", valorFinal)
}

func recibirEnteros(i string, x bool, n ...int) {
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Println(i)
	fmt.Printf("%T\n", i)
}
func sumaNumeros1(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	suma := 0
	for _, v := range x {
		suma += v
		//suma = suma + v
	}
	fmt.Println(suma)
	return suma

}
