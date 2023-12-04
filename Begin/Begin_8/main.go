package main

import "fmt"

func main () {
	var a, b float32
	fmt.Println("Enter the number for a:")
	fmt.Scan(&a)

	fmt.Println("Enter the number for b:")
	fmt.Scan(&b)
	var Sr = (a + b) / 2
	
fmt.Println("Arithmetic mean of numbers:",Sr)

}