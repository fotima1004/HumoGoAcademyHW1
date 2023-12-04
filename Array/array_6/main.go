package main

import "fmt"

func main()  {
	var n int 
	fmt.Scan(&n)
	array := make([] int,n)
	for i := 0; i < n; i++{
		fmt.Scan(&array[i])	
	}
	min := array[0]
	var inmin int = 0
	for i := 1; i < n; i++ {
		if array[i] < min{
			min = array[i]
			inmin = i

		}
		
	}
	array[inmin]++
	res := 1
	for i := 0; i < n; i++ {
		res *= array[i]
		
	}
	println(res)

}