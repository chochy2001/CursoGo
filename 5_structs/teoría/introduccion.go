package main

import "fmt"

/*
type persona struct {
	nombre,apellido   string
	edad     uint8
	estatura float32
	sexo     bool
}

type estudiante struct{
	persona
	curso string
	grupo int
}
*/

func main() {
	persona1 := persona{
		nombre:   "Juan",
		edad:     25,
		estatura: 1.75,
	}
	persona2 := persona{
		"Pedro",
		"Gonzalez",
		30,
		2.2,
		true,
	}

	persona3 := persona{
		apellido: "Lopez",
		estatura: 1.75,
		nombre:   "Maria",
		edad:     25,
	}
	/*
		persona4:= persona{
			"Lopez",
			1.75,
			"Maria",
			25,
		}
	*/
	fmt.Println(persona1)
	fmt.Println(persona1.nombre, persona1.apellido, persona1.edad, persona1.estatura)
	fmt.Printf("Tu nombre es: %s %s, tienes %d años y tu estatura es de %.03f\n", persona1.nombre, persona1.apellido, persona1.edad, persona1.estatura)
	fmt.Printf("Tu nombre es: %v %v, tienes %v años y tu estatura es de %v\n", persona1.nombre, persona1.apellido, persona1.edad, persona1.estatura)
	fmt.Println("Tu nombre es: ", persona1.nombre, persona1.apellido, "tienes", persona1.edad, "años y tu estatura es de", persona1.estatura)

	fmt.Println(persona2)
	fmt.Println(persona3)

	//Una estructura es una secuencia de elementos con nombre, llamados campos,
	//cada uno de los cuales tiene un nombre y un tipo. Los nombres de los campos pueden
	//especificarse explícitamente (IdentifierList) o implícitamente (EmbeddedField).
	//Dentro de una estructura, los nombres de campo que no estén en blanco deben ser únicos.

}
