package main

import "fmt"

func main () {
	var a, b int
	fmt.Println("Enter the number for a:")
	fmt.Scan(&a)

	fmt.Println("Enter the number for b:")
	fmt.Scan(&b)

	var S = a * b
	var P = 2 * (a + b)
	
fmt.Println("Area of the rectangle:",S)
fmt.Println()
fmt.Println("Perimetr of the rectangle:",P)
}