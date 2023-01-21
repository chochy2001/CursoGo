package main

import "fmt"

//Crea un vehículo con una estructura que contenga datos como el
//color,
//marca, modelo y número de puertas.

//Genera dos tipos de vehiculos como son camioneta y automovil (carro si lo prefieres)

//Embebe vehículo en ambos tipos de estructuras (camioneta y automovil)

//Agrega a cada uno de las estructuras (camioneta y automovil)
//los campos de "lujo" y "cuatroPuertas". (como booleanos)

//Crea los 2 tipos de vehiculos y asignales sus valores.

type vehiculo struct {
	color   string
	marca   string
	modelo  int
	puertas int
}

type camioneta struct {
	vehiculo
	lujo bool
}

type carro struct {
	vehiculo
	cuatroPuertas bool
}

func main() {

	myCar := carro{
		cuatroPuertas: true,
		vehiculo: vehiculo{
			"negro",
			"ford",
			2023,
			4,
		},
	}

	myTruck := camioneta{
		lujo: false,
		vehiculo: vehiculo{
			color: "blanco",
			marca: "chevrolet",
			//modelo:  2025,
			puertas: 4,
		},
	}

	fmt.Println(myCar, myTruck)

}
