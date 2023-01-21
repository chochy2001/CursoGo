package main

import "fmt"

func main() {
	persona := struct {
		//atributos
		nombre         string
		apellidos      string
		edad           int
		contactos      map[string]int
		comidaFavorita []string
	}{
		//inicializacion
		nombre:    "Jorge",
		apellidos: "Salgado",
		edad:      21,
		contactos: map[string]int{
			"Juan":   1234,
			"Andrea": 2345,
			"Pedro":  34538,
		},
		comidaFavorita: []string{
			"enchiladas",
			"mole",
			"pozole",
			"Pizza",
		},
	}
	fmt.Println(persona)
	fmt.Println(persona.contactos["Andrea"])
	for i, v := range persona.contactos {
		fmt.Println(i, v)
	}
	for i, v := range persona.comidaFavorita {
		fmt.Println(i, v)
	}

}
