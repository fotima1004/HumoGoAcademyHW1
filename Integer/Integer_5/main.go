package main

import "fmt"

func main () {
	var A, B int
	fmt.Println("Enter the length of segment A:")
	fmt.Scan(&A)

	fmt.Println("Enter the length of segment B:")
	fmt.Scan(&B)
	if A > B {
		var Result = A % B
	fmt.Println("Segment A contains ",Result)	
	}else {
		fmt.Println("Does not satisfy the condition.")
	}
}	