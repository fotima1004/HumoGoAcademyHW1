package main

import "fmt"
func main () {
	var a int 
	fmt.Println("Enter the number for a:")
	fmt.Scan(&a)
	if a > 0{
		fmt.Println("number a is positive: ",a + 1)		
	}else {
		fmt.Println("number a is not positive: ",a - 2)	
	}
}