package main

import "fmt"

type personaMejorado struct {
	nombres   []string
	apellidos []string
	edad      uint8
	peso      float32
	estatura  float32
	hobbies   []string
}

func main() {
	//Crea una estructura llamada persona con los siguientes campos:
	//nombre, apellido, edad, peso, estatura, y hobbie.
	//Crea dos personas y muestra sus datos. con un formato especifico

	// * Hola, mi nombre es: Juan, mi apellido es: Pérez, tengo: 23 años, peso: 65 kg,
	// * estatura: 1.78 m y mi hobbie es: fútbol.
	persona1 := personaMejorado{
		nombres:   []string{"Juan", "Pedro"},
		apellidos: []string{"Pérez", "Gómez"},
		edad:      23,
		peso:      65,
		estatura:  1.78,
		hobbies:   []string{"fútbol"},
	}
	persona2 := personaMejorado{
		nombres:   []string{"Jorge"},
		apellidos: []string{"Salgado", "Miranda"},
		edad:      20,
		peso:      100,
		estatura:  1.85,
		hobbies:   []string{"fútbol", "natación", "caminar", "programar"},
	}

	//fmt.Println(persona1,persona2)
	fmt.Println("Hola, mi nombre es:", persona1.nombres, "mi apellido es:", persona1.apellidos, "tengo:", persona1.edad, "años, \npeso:", persona1.peso, "kg, estatura:", persona1.estatura, "m y mi hobbie es:", persona1.hobbies)
	fmt.Println("Hola, mi nombre es:", persona2.nombres, "mi apellido es:", persona2.apellidos, "tengo:", persona2.edad, "años, \npeso:", persona2.peso, "kg, estatura:", persona2.estatura, "m y mi hobbie es:", persona2.hobbies)

	for i, v := range persona1.nombres {
		fmt.Println(i, v)
	}
	for i, v := range persona1.apellidos {
		fmt.Println(i, v)
	}
	for i, v := range persona1.hobbies {
		fmt.Println(i, v)
	}

}
