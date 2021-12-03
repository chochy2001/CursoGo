package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	fmt.Println(x)
	y := []string{"e", "f", "g", "h"}

	fmt.Println(y)

	xy := [][]string{x, y}
	/*
		fmt.Println(xy)
		xy[0][0] = "x"
		fmt.Println(xy)
		xy[1][2] = "w"
		fmt.Println(xy)

	*/

	for i := 0; i < 2; i++ {
		fmt.Println(i)
		for j := 0; j < 4; j++ {
			fmt.Println(j)
			fmt.Println(xy[i][j])
		}
	}

}
