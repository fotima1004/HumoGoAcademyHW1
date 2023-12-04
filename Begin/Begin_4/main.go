package main

import "fmt"

func main () {
	var d float32 
	const pi = 3.14 
	fmt.Println("Enter the circle diameter number:")
	fmt.Scan(&d)
	var L =  pi * d	
fmt.Println("Circumference:",L)
}