package main

import "fmt"

func main () {
	var n int
	fmt.Println("Enter a three-digit number:")
	fmt.Scan(&n)
	var res = n % 10
	var result = n / 10 % 10

	fmt.Println("Last digit:",res)
	fmt.Println("Average digit:",result)
}	