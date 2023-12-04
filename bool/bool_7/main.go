package main

import "fmt"
func main () {
	var a, b, c int 
	fmt.Println("Enter the number for a:")
	fmt.Scan(&a)
	fmt.Println("Enter the number: for b:")
	fmt.Scan(&b)
	fmt.Println("Enter the number: for c:")
	fmt.Scan(&c)
	var res = b > a && b < c
	
fmt.Println(res)
}