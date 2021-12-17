package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     uint8
	peso     float32
	estatura float32
	hobbie   string
}

func main() {
	//Crea una estructura llamada persona con los siguientes campos:
	//nombre, apellido, edad, peso, estatura, y hobbie.
	//Crea dos personas y muestra sus datos. con un formato especifico

	// * Hola, mi nombre es: Juan, mi apellido es: Pérez, tengo: 23 años, peso: 65 kg,
	// * estatura: 1.78 m y mi hobbie es: fútbol.
	persona1 := persona{
		nombre:   "Juan",
		apellido: "Pérez",
		edad:     23,
		peso:     65,
		estatura: 1.78,
		hobbie:   "fútbol",
	}
	persona2 := persona{
		nombre:   "Jorge",
		apellido: "Salgado",
		edad:     20,
		peso:     100,
		estatura: 1.85,
		hobbie:   "fútbol",
	}

	fmt.Println(persona1, persona2)
	fmt.Println("Hola, mi nombre es:", persona1.nombre, "mi apellido es:", persona1.apellido, "tengo:", persona1.edad, "años, peso:", persona1.peso, "kg, estatura:", persona1.estatura, "m y mi hobbie es:", persona1.hobbie)
	fmt.Println("Hola, mi nombre es:", persona2.nombre, "mi apellido es:", persona2.apellido, "tengo:", persona2.edad, "años, peso:", persona2.peso, "kg, estatura:", persona2.estatura, "m y mi hobbie es:", persona2.hobbie)

}
