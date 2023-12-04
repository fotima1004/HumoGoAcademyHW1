package main

import "fmt"

func main () {
	var n int
	fmt.Println("Enter a two-digit number:")
	fmt.Scan(&n)
	var res = (n % 10) * 10 + (n / 10)
	fmt.Println("Number obtained by rearrangement",res)
}	