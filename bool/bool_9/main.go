package main

import "fmt"
func main () {
	var a, b int 
	fmt.Println("Enter the number for a:")
	fmt.Scan(&a)
	fmt.Println("Enter the number: for b:")
	fmt.Scan(&b)
	var res = a % 2 != 0 || b % 2 != 0

fmt.Println(res)
}