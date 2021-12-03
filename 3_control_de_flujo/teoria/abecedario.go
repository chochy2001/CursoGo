package main

import "fmt"

func main(){
	//Imprimir el abecedario en mayusculas y en minusculas
	//a
	//b
	//c
	//a b c

	//letra A = 65
	//letra Z = 90
	fmt.Println("Abecedario en Mayusculas A-Z")
	for i:=65;i<=90; i++{
		fmt.Printf("%c ",i)
	}


	//letra a = 97
	//letra z = 122
	fmt.Println("\nAbecedario en Minúsculas a-z")
	for i:=97;i<=122; i++{
		fmt.Printf("%c ",i)
	}





	//letra A = 65
	//letra Z = 90
	fmt.Println("\nAbecedario en Mayusculas Z-A")
	for i:=90;i>=65; i--{
		fmt.Printf("%c ",i)
	}


	//letra a = 97
	//letra z = 122
	fmt.Println("\nAbecedario en Minúsculas z-a")
	for i:=122;i>=97; i--{
		fmt.Printf("%c ",i)
	}



	fmt.Println("\n Abecedario sin las vocales")
	//a = 97
	//e = 101
	//i = 105
	//o = 111
	//u = 117

	for i:=122;i>=97; i--{
		if i == 97 || i == 101 || i == 105 || i == 111 || i == 117{
			continue
		}
		fmt.Printf("%c ",i)
	}







}
