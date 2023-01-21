package main

import "fmt"

//Apuntadores
//& Dirección en memoria de la variable

func main() {
	variable := 21.2
	fmt.Println(variable)
	fmt.Println(&variable)

	fmt.Printf("%T\n", variable)
	fmt.Printf("%T\n", &variable)

	var variableAlmacenamiento *float64 = &variable
	fmt.Printf("Tipo dato variableAlmacenamiento: %T\n", variableAlmacenamiento)
	fmt.Printf("Valor &variableAlmacenamiento: %v\n", &variableAlmacenamiento)
	fmt.Printf("Valor &variable: %v\n", &variable)
	fmt.Printf("Valor variableAlmacenamiento: %v\n", variableAlmacenamiento)
	//Todo en Go se pasa por valor

	*variableAlmacenamiento = 10.2
	fmt.Printf("\n\nvalor variableAlmacenamiento (Reasignado): %v\n", &variableAlmacenamiento)
	fmt.Printf("valor variable: %v\n", variable)
	fmt.Printf("direccion memoria variable: %v\n", &variable)
	fmt.Printf("valor en variable apuntada variableAlmacenamiento: %v\n", *variableAlmacenamiento)
	fmt.Printf("dirección en memoria del valor en variable apuntada variableAlmacenamiento: %v\n", *&variableAlmacenamiento)

}
