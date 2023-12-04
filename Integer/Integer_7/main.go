package main


import "fmt"

func main () {
	var n int
	fmt.Println("Enter a two-digit number:")
	fmt.Scan(&n)
	var sum = (n/10) + (n % 10);

	var mult = (n/10) * (n % 10);
	
	fmt.Println("Sum of two numbers =",sum)
	fmt.Println("Product of two numbers =",mult)
	
}	