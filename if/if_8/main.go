package main

import "fmt"
func main () {
	var a int 
	fmt.Println("Enter the number for a:")
	fmt.Scan(&a)
	var b int 
	fmt.Println("Enter the number for b:")
	fmt.Scan(&b)
	if a > b{
		fmt.Println(a,"is big")
		fmt.Println(b,"is small")
	}else if b > a{
		fmt.Println(b,"is big")
		fmt.Println(a,"is small")
	}else {
		fmt.Println(a,"=",b,"a and b are equal")	
	}		
}