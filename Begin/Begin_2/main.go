package main

import ("fmt"
		"math"
	)

func main () {
	var a float64 
	fmt.Println("Enter the number:")
	fmt.Scan(&a)
	var S =  math.Pow(a, 2)	
fmt.Println("Area of the square:",S)
}