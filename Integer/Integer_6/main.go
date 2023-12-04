package main

import "fmt"

func main () {
	var n int
	fmt.Println("Enter a two-digit number:")
	fmt.Scan(&n)
	
	fmt.Println("Left digit",n / 10)
	fmt.Println("Right digit",n % 10)
	
}	