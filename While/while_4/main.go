package main

import (
	"fmt"
)

func main()  {
	var n int
	fmt.Println("Enter the number for n:")
	fmt.Scan(&n)
	var k int = 1
	for k < n{
		k *= 3
	}
	fmt.Println(k==n)
}