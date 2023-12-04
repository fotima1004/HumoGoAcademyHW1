package main

import "fmt"

func main () {
	var a int
	fmt.Println("Enter the number:")
	fmt.Scan(&a)
	var result = a % 2 == 0
	
fmt.Println(result)
}