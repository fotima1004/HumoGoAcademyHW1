package main

import (
	"fmt"
	"math"
)

func main () {

	var a, b float64

	fmt.Println("Enter the number for a:")
	fmt.Scan(&a)
	fmt.Println("Enter the number for b:")
	fmt.Scan(&b)
	if (a > 0 && b > 0) {
		var k = math.Sqrt(a * b)
		fmt.Println("Geometric mean of numbers:",k)
	}else {
		fmt.Println("Numbers do not meet requirements:")
	}
}