package main

import "fmt"
func main () {
	var a int 
	fmt.Println("Enter the number for a:")
	fmt.Scan(&a)
	if a > 0{
		fmt.Println("number a is positive: ",a + 1)		
	}else if a < 0{
		fmt.Println("number a is not positive: ",a - 2)	
	}else {
		fmt.Println("the number a is equal to zero: ",a + 10)	
	}
}