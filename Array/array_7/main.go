package main

import "fmt"

func main()  {
	var Array_Size, Quantity  int 
	fmt.Println("Enter the number of elements in the array: N")
	fmt.Scan(&Array_Size)
	array := make([] int,Array_Size)
	for i := 0; i < Array_Size; i++{
		fmt.Scan(&array[i])	
	}
	var x int
	fmt.Println("Enter the number for X")
	fmt.Scan(&x)
	for i := 0; i < Array_Size; i++{
		if array[i] == x{	
			Quantity ++
		}
	}
	fmt.Println("Quantity:",Quantity)
}