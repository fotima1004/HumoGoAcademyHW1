package main

import "fmt"

func main () {
	var a, b, c float64
	fmt.Println("Enter the number for a:")
	fmt.Scan(&a)
	fmt.Println("Enter the number for b:")
	fmt.Scan(&b)
	fmt.Println("Enter the number for c:")
	fmt.Scan(&c)
	var V  = a * b * c
	var S = 2 * ((a * b) + (b * c) + (a * c))
fmt.Println("Volume of a rectangular parallelepipeda:",V)
fmt.Println("Surface area:",S)
}