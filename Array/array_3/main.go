package main

import "fmt"

func main()  {
	var n int 
	fmt.Println("Enter the number of elements in the array: N")
	fmt.Scan(&n)
	array := make([] int,n)
	for i := 0; i < n; i++{
		fmt.Scan(&array[i])	
	}
	var max = array[0]
	for i := 1; i < n; i++{
		if array[i] > max{	
			max = array[i]
		}
	}
	fmt.Println("Maximum:",max)
}