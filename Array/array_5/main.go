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
	var minOdd int 
	var found bool
	for i :=0; i < n; i++{
		 // min 
		 if (array[i] % 2 != 0) && (array[i] <minOdd || !found){
			minOdd = array[i]
			found = true
		}
		}
		
	if found {
		fmt.Println(minOdd)
	} else {
		fmt.Println(0)
	}
}