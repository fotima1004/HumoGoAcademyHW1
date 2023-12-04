package main

import "fmt"
func main () {
	var a, b int 
	fmt.Println("Enter the number for a:   a >= 0")
	fmt.Scan(&a)
	fmt.Println("Enter the number: for b:  b < -2")
	fmt.Scan(&b)

fmt.Println(a >= 0 || b < -2)
}