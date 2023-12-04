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
	for i := 0; i < n / 2; i++ {
		array[i],array[n-i-1] = array[n - i - 1], array[i]
		
	}
	fmt.Println(array)









	/*var x int
	for i := 0; i < n / 2; i ++{ 
		x = array[n - i - 1]
		array[n - i - 1] = array[i]
		array[i] = x
	}
	fmt.Println(array)
	*/


}