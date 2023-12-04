package main

import "fmt"

func main () {
	var a, b int 
	fmt.Println("Enter the number for a:  a > 2")
	fmt.Scan(&a)
	fmt.Println("Enter the number: for b:  b <= 3")
	fmt.Scan(&b)
	
fmt.Println(a > 2 && b <= 3)
}