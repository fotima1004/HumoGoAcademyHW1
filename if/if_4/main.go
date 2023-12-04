package main

import "fmt"
func main () {
	var quantity int = 0
	var a int 
	fmt.Println("Enter the number for a:")
	fmt.Scan(&a)
	var b int 
	fmt.Println("Enter the number for b:")
	fmt.Scan(&b)
	var c int 
	fmt.Println("Enter the number for c:")
	fmt.Scan(&c)
	if a > 0{
		quantity += 1
	}
	if b > 0{
		quantity += 1
	}
	if c > 0{
		quantity += 1
	}			
	fmt.Println("number of positive",quantity)			
}