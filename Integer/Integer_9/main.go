package main

import "fmt"

func main () {
	var n int
	fmt.Println("Enter a three-digit number:")
	fmt.Scan(&n)
	if n >= 100 {
		var res = n / 100 
	fmt.Println("res =",res)	
	}else {
		fmt.Println("Does not satisfy the condition.")
	}
}	