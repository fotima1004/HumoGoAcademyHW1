package main

import ("fmt"
		"math"
		)

func main () {
	var a float64 

	fmt.Println("Enter the length of the cube edge:")
	fmt.Scan(&a)
	var V = math.Pow(a,3)
	var S = 6 * math.Pow(a,2)
fmt.Println("Cube volme:",V)
fmt.Println("Surface area:",S)

}