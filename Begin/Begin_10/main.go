package main

import
	"fmt"

func main () {

	var n, k float64

	fmt.Println("Enter the number for a:")
	fmt.Scan(&n)
	fmt.Println("Enter the number for b:")
	fmt.Scan(&k)
	if (n > 0 && k > 0) {

		var sum = n + k
		fmt.Println("Sum of two numbers:",sum)
	
		var diff = n - k
		fmt.Println("difference of two numbers:",diff)
		
		var pro = n * k
		fmt.Println("product of two numbers:",pro)
		
		var quot = n / k
		fmt.Println("quotient of their squares:",quot)
	}else {
		fmt.Println("Numbers do not meet requirements:")		
	}
}