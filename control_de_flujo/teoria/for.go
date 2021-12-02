// El bucle Go for es similar -pero no igual- al de C.
// Unifica for y while y no hay do-while. Hay tres formas, de las cuales sólo una tiene punto y coma.

/*
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }
 */

package main

import "fmt"

func main (){
	/*
	//Imprimir en pantalla los numeros del 1 al 100
	fmt.Println("Imprime los valores desde el 1 hasta el 100 sumando de 1 en 1")
	for i := 1; i<= 100; i++{
		fmt.Println(i)
	}
	// Imprimir en pantalla los numeros del 100 al 1
	fmt.Println("Imprime los valores desde el 100 hasta el 1 restando de 1 en 1")
	for i := 100; i>= 1; i--{
		fmt.Println(i)
	}


	//Imprime los valores del 1 al 100 sumando de 2 en 2
	for i:= 1; i<= 100;i+=2{
		fmt.Println(i)
	}
	//Imprime los valores del 1 al 100 sumando de 5 en 5
	for i:= 1; i<= 100;i+=5{
		fmt.Println(i)
	}
	//Imprime los valores del 100 al 1 restando de 2 en 2
	for i:= 100; i>= 0;i-=2{
		fmt.Println(i)
	}

	//Imprime los valores del 100 al 1 restando de 5 en 5
	for i:= 100; i>= 0;i-=5{
		fmt.Println(i)
	}
	//Imprime todos los años que has estado en vivo
	fmt.Println("Estos son todos los años que he estado vivo")
	for i:= 2001; i<= 2021;i++{
		fmt.Println(i)
	}

			a := 19
	for a >= 18{
		fmt.Println("Eres mayor de edad")
		a--
	}

	for i:= 0; i <= 5; i++ {
		fmt.Printf("Bucle externo: %d \n", i)
		for j:= 0; j <= 3; j++ {
			fmt.Printf("\tBucle interno: %d \n", j)
		}
	}
	*/


	for i:= 0; i <= 5; i++ {
		fmt.Println(i)
		if i == 3{
			continue
		}
		fmt.Println("Esto es una prueba")
		//
		//
		//
		//
		//

	}






}


