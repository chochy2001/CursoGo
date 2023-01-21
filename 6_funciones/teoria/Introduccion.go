package main

import "fmt"

func main() {
	hola()
	holaTexto("Jorge")

	adios()
	holaTexto("Juan")
	adiosNumero(10)

	fmt.Println(sumaNumeros(3, 4))
	fmt.Println(3)
	sumaTotal := sumaNumeros(1, 2)
	fmt.Println(sumaTotal)

}

func sumaNumeros(numero1 int, numero2 int) int {
	return numero1 + numero2
}
func hola() {
	fmt.Println("Hola desde funcion hola")
}
func holaTexto(texto string) {
	fmt.Println("Hola,", texto)
}
func adios() {
	fmt.Println("Adios desde funcion adios")
}
func adiosNumero(numero int) {
	fmt.Println("Hola, tu numero es: ", numero)
}
