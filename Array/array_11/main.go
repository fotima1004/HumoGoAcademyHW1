package main

import "fmt"

func main()  {
	var Array_Size int 
	fmt.Println("Enter the number of elements in the array: N")
	fmt.Scan(&Array_Size)
	array := make([] int,Array_Size)
	for i := 0; i < Array_Size; i++{
		fmt.Scan(&array[i])	
	}
	max := array[0]
	min := array[0]
	for i := 1; i < Array_Size; i++{
		if array[i] > max{
			max = array[i] 
		}
		if array[i] < min{
			min = array[i] 
		}
	}
	fmt.Println("After processing the array")
	for i := 0; i < Array_Size; i++ {
		if array[i] == max{
			array[i] = min
		}
		println(array[i])	
	}

}