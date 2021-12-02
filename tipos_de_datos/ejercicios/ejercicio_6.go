// * Usando Iota crea n constantes desde tu año de nacimiento hasta el año actual
// * (limitalo a 5 años si lo prefieres).
// ! o usa un ciclo para recorrer todos los años
// * Imprime los valores de estas constantes

// ? Explica que es lo que hace Iota y en que casos se puede usar

package main

import "fmt"

const(
	year = 2001 + iota
	year1
	year2
	year3
	year4
	year5
	year6
	year7
	year8
	year9
	year10
)

func main (){
	fmt.Println(year, year1, year2, year3, year4,year5,year6,year7,year8,year9,year10)
}
