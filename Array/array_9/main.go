package main

import(
	"fmt"
	"math"
)

func main()  {
	var Array_Size int 
	fmt.Println("Enter the number of elements in the array: N")
	fmt.Scan(&Array_Size)
	array := make([] int,Array_Size)
	for i := 0; i < Array_Size; i++{
		fmt.Scan(&array[i])	
	}
	var x int
	min := array[0]
	fmt.Println("Enter the number for X")
	fmt.Scan(&x)
	for i := 0; i < Array_Size; i++{
		if math.Abs(float64(array[i]-x)) < math.Abs(float64(min - x))||
		 float64((array[i]-x)) == math.Abs(float64(min - x)) && array[i] < min{
			min = array[i]
		}	
	}
	fmt.Println(min)
}