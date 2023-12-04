package main

import "fmt"
func main () {
	var res float64
	var a float64 
	fmt.Println("Enter the number for a:")
	fmt.Scan(&a)
	var b float64 
	fmt.Println("Enter the number for b:")
	fmt.Scan(&b)
	if a < b{
		fmt.Println(a,"is less ",b)
		fmt.Println(b,"is big ",a)
	}else if b < a{
		res = a
		a = b
		b = res
		fmt.Println("new value a =",a)
		fmt.Println("new value b =",b)
	}else {
		fmt.Println("a and b are equal:", a,"=",b)
	}
}