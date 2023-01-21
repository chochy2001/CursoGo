package main

import "fmt"

//Crea una estructura con identificador auto
//Que contenga año(modelo), dueño, puertas
//Agregale el metodo manejar
//(el cual tiene que tener un texto diciendo que se esta manejando junto con su modelo)

//Crea un nuevo auto usando la estructura con identificador ferrari
//y llama al metodo manejar del ferrari

type auto struct {
	modelo, puertas int
	duenio          string
}

func (a auto) manejar() {
	fmt.Printf("¡¡Manejando!!\nmodelo: %v,\ndueño: %v,\npuertas: %v,\n", a.modelo, a.duenio, a.puertas)
}

func main() {
	ferrari := auto{
		modelo:  2025,
		puertas: 2,
		duenio:  "Jorge Salgado",
	}
	ferrari.manejar()
}
