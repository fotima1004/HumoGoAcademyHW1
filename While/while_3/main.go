package main

import "fmt"

func main()  {
	var n int
	fmt.Println("Enter the number for n:")
	fmt.Scan(&n)
	var k int
	fmt.Println("Enter the number for k:")
	fmt.Scan(&k)
	var Kol int 
	for n > k{
		n -= k
		Kol += 1
	}
	fmt.Println("Quotient of dividing N by K:",Kol)
	fmt.Println("Modulo:",n)
}
