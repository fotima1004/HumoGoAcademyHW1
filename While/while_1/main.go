package main

import "fmt"

func main()  {
	var a int
	fmt.Println("Enter the number for a:")
	fmt.Scan(&a)
	var b int
	fmt.Println("Enter the number for b:")
	fmt.Scan(&b)
	for a > b{
		a -= b
	}
	fmt.Println("Length of the unoccupied part of segment a:",a)
}