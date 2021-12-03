package main

import "fmt"

func main() {
	//Los mapas son una cómoda y potente estructura de datos integrada que asocia valores de un tipo (la clave)
	//con valores de otro tipo (el elemento o el valor).

	//La clave puede ser de cualquier tipo para el que esté definido el operador de igualdad, como enteros,
	//números de coma flotante y complejos, cadenas, punteros, interfaces (siempre que el tipo dinámico
	//admita la igualdad), structs y arrays. Las rebanadas no pueden utilizarse como claves de mapa,
	//porque la igualdad no está definida en ellas. Al igual que los slices, los mapas contienen referencias
	//a una estructura de datos subyacente. Si se pasa un mapa a una función que cambia el contenido del mapa,
	//los cambios serán visibles en el llamador.
	//
	//Los mapas pueden construirse utilizando la sintaxis literal compuesta habitual con pares clave-valor
	//separados por dos puntos, por lo que es fácil construirlos durante la inicialización.
	mapa := map[string]int{
		"Juan":  5512323455,
		"Pedro": 6546546546,
		"Luis":  7657657657,
		"Maria": 8888888888,
		"Ana":   9999999999,
		"Jorge": 1234560,
	}
	maps := map[string]string{
		"Juan":  "5512323455",
		"Pedro": "6546546546",
		"Luis":  "7657657657",
		"Maria": "8888888888",
		"Ana":   "9999999999",
	}
	fmt.Println("Mapa:", mapa)
	fmt.Println(len(mapa))
	fmt.Println("Maps:", maps)
	fmt.Println(len(maps))

	fmt.Println(mapa["Juan"])
	fmt.Printf("%T\n", mapa["Juan"])
	fmt.Println(maps["Juan"])
	fmt.Printf("%T\n", maps["Juan"])
	fmt.Println(maps["Jorge"])
	fmt.Println(mapa["Jorge"])

	/*_, ok := mapa["Jorge"]
	fmt.Println(ok)
	*/

	if v, ok := mapa["Ana"]; ok {
		fmt.Println("El numero celular de Ana es:", v)
	} else {
		fmt.Println("No existe")
	}

}
