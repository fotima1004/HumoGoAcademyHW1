package main

import "fmt"
func main () {
	var a int 
	fmt.Println("Enter the number for a:")
	fmt.Scan(&a)
	var b int 
	fmt.Println("Enter the number for b:")
	fmt.Scan(&b)
	if a < b{
		fmt.Println(a,"is less than",b)
	}else if b < a{
		fmt.Println(b,"is less than",a)	
	}else {
		fmt.Println(a,"=",b)
	}		
}