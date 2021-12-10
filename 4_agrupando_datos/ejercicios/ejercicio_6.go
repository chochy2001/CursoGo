package main

import "fmt"

func main() {
	//crea un slice con nombre de alumnos, normalmente los grupos son de 5 personas,
	//pero luego se inscriben más personas, prevé eso. el máximo puede ser de 10
	//que el nombre de los alumnos los ingrese el usuario
	// ? Cual seria la capacidad del slice -> 10
	// ? Cual seria la longitud del slice -> 5
	// ? Imprime el contenido del slice con sus indices y su valor

	nombreAlumnos := make([]string, 5, 10)
	for i, _ := range nombreAlumnos {
		fmt.Println("Ingresa el nombre del alumno: ")
		//fmt.Scanf("%s",&nombreAlumnos[i])
		fmt.Scanln(&nombreAlumnos[i])
		//fmt.Scan(&nombreAlumnos[i])
	}
	fmt.Println(nombreAlumnos)
	for i, valor := range nombreAlumnos {
		fmt.Println("El nombre del alumno en la posicion ", i, " es: ", valor)
	}
	fmt.Println("\n")

	for i := 0; i < len(nombreAlumnos); i++ {
		fmt.Println("El nombre del alumno en la posicion ", i, " es: ", nombreAlumnos[i])
	}

}
