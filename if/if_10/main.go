package main

import "fmt"
func main () {
	var res int
	var a int 
	fmt.Println("Enter the number for a:")
	fmt.Scan(&a)
	var b int 
	fmt.Println("Enter the number for b:")
	fmt.Scan(&b)
	if a != b {
		res = a + b
		a = res
		b = res
		fmt.Println("new value a =",a)
		fmt.Println("new value b =",b)
	}else {
		a = 0
		b = 0
		fmt.Println("a =",a)
		fmt.Println("b =",b)
	}
}