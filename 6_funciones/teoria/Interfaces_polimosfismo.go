package main

import "fmt"

//Las interfaces sirven para crear polimorfismo y especificar un comportamiento

type persona struct {
	nombre string
	edad   uint8
}

type estudiante struct {
	persona
	id int
}

// func (identificador estructura) identificadorFuncion(parametros) (returns){
// codigo}

func (e estudiante) hablar() {
	fmt.Println("Mi nombre es:", e.nombre, " mi id:", e.id, "mi edad:", e.edad)
}
func (p persona) hablar() {
	fmt.Println("La persona esta hablando")
}
func (p persona) caminar() {
	fmt.Println("Estoy caminando")
}
func (p persona) respirar() {
	fmt.Println("Estoy respirando")
}

type serVivo interface {
	respirar()
}

func xy(s serVivo) {
	switch s.(type) {
	case persona:
		fmt.Println("Soy una persona", s.(persona).edad, s.(persona).nombre)
	case estudiante:
		fmt.Println("Soy un estudiante", s.(estudiante).edad, s.(estudiante).id, s.(estudiante).nombre, s.(estudiante).persona)
	}
}

func main() {
	est1 := estudiante{
		persona: persona{
			nombre: "Jorge",
			edad:   21,
		},
		id: 12340,
	}
	est2 := estudiante{
		persona: persona{
			nombre: "Ricardo",
			edad:   20,
		},
		id: 2378745,
	}
	per1 := persona{
		nombre: "Juan",
		edad:   10,
	}
	est1.hablar()
	per1.hablar()

	xy(est1)
	xy(est2)
	xy(per1)

	/*
		per1.caminar()
		est1.hablar()
		est2.hablar()
		per1.hablar()
		est1.hablar()
		est2.hablar()
	*/

}
