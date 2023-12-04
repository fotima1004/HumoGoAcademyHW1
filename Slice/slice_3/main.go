package main

import
	"fmt"


func main()  {
	var n int
	fmt.Scan(&n)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])	
	}
	var k, j int
	
	fmt.Scan(&k)
	for i := 0; i < n; i++{
		fmt.Print(array[i]," ")
		j ++
		if j == k{
			j = 0
			fmt.Println()	
		}

	}

}
