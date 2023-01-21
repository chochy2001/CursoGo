package main

import "fmt"

func main() {
	//nombre()
	//nombre1(21)
	funcion1 := func() {
		fmt.Println("Hola desde función anónima")
	}
	funcion1()

	funcion2 := func() bool {
		return true
	}()
	//funcion2()

	fmt.Println(funcion2)

	func(edad uint8) {
		fmt.Println("Hola, mi edad es ", edad, "desde función anónima")
	}(21)

}

func nombre() {
	fmt.Println("Hola")
}
func nombre1(edad uint8) {
	fmt.Println("Hola, mi edad es ", edad)
}
