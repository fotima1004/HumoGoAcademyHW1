package main

import (
	"fmt"
)

func main()  {
	var n int
	fmt.Println("Enter the number for N:")
	fmt.Scan(&n)
	var m int 
	var factorial int = 1
	for m < n{
	factorial *= (n - m)
	m += 2
	}
	fmt.Println("Double factorial N:",factorial)
}