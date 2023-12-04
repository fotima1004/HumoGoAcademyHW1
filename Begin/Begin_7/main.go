package main

import (
	"fmt"
	"math"
)
func main () {
	var  R float64
	const pi = 3.14
	fmt.Println("Enter radius:")
	fmt.Scan(&R)

	var L = 2 * pi * R
	var S = pi * math.Pow(R,2)
	 
fmt.Println("Circumference",L)
fmt.Println("Area of a circle",S)
}
