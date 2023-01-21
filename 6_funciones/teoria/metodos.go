package main

//Metodos

//Funciones dentro de una clase (En otros Lenguajes)

/*
type persona struct {
	nombre string
	edad   uint8
}

type estudiante struct {
	persona
	id int
}
*/

// func (identificador estructura) identificadorFuncion(parametros) (returns){
// codigo}

/*
func (e estudiante) hablar() {
	fmt.Println("Mi nombre es:", e.nombre, " mi id:", e.id, "mi edad:", e.edad)
}
func (p persona) caminar() {
	fmt.Println("Estoy caminando")
}

*/

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
	per1.caminar()
	est1.hablar()
	est2.hablar()

}
