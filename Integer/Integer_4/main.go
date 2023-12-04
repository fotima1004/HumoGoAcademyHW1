package main

import "fmt"

func main () {
	var A, B int
	fmt.Println("Enter the length of segment A:")
	fmt.Scan(&A)

	fmt.Println("Enter the length of segment B:")
	fmt.Scan(&B)
	if A > B {
		var Res = A / B
		fmt.Println("Segment A contains ",Res," of the complete segments of B.")	
	}else {
		fmt.Println("Does not satisfy the condition.")
	}
}	