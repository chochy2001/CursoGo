package main

import "fmt"

//Crea y usa una función anonima para lo que quieras

func main() {
	//Función anónima sin parametros
	func() {
		fmt.Println("Hola")
	}()

	//Función anónima con parametros
	func(edad int) {
		fmt.Println("Hola, mi edad es: ", edad)
	}(21)

	//Función anónima asignada a una variable
	xi := func() {
		fmt.Println("Hola")
	}
	xi()
	//Función anónima dentro de otra
	func() {
		fmt.Println("Hola fuera")
		func() {
			fmt.Println("Hola, mi dentro de otra")
		}()
	}()
	//Función anónima dentro de otra asignado a una variable
	y1 := func() {
		fmt.Println("Hola fuera")
		func() {
			fmt.Println("Hola, mi dentro de otra")
		}()
	}
	y1()

}
