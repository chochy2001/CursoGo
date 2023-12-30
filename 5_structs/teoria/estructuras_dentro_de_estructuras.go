package main

import "fmt"

type persona struct {
	nombre, apellido string
	edad             uint8
	estatura         float32
	sexo             bool
}

type estudiante struct {
	persona
	curso string
	grupo int
}

func main() {
	estudiante1 := estudiante{
		persona: persona{
			nombre:   "Juan",
			apellido: "Perez",
			edad:     20,
			estatura: 1.80,
			sexo:     true,
		},
		curso: "Ingenieria de Software",
		grupo: 1,
	}
	fmt.Println(estudiante1)
	fmt.Println(estudiante1.persona.nombre)
	fmt.Println(estudiante1.nombre)
	//fmt.Println(estudiante1.persona.apellido)
	//fmt.Println(estudiante1.edad)

}
