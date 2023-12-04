package main

import (
	"fmt"
)

func main()  {
	var n int
	fmt.Println("Enter the number for n:")
	fmt.Scan(&n)
	var m int = 1 
	var k int
	for m <= n{
		m *= 2
		k ++
	}
	fmt.Println("Indicator of this degree K:",k)
}

