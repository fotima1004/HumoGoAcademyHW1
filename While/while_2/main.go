package main

import "fmt"

func main()  {
	var a int
	fmt.Println("Enter the number for a:")
	fmt.Scan(&a)
	var b int
	fmt.Println("Enter the number for b:")
	fmt.Scan(&b)
	var Kol int 
	for a > b{
		a -= b
		Kol += 1
	}
	fmt.Println("Number of segment b placed on segmenta a:",Kol)
}
