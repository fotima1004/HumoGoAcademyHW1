package main

import(
	"fmt"
	"math"
)


func main()  {
	var n int
	fmt.Scan(&n)
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])	
	}
	fmt.Println(array)
	minPositive := math.MaxInt
	maxNegative := math.MinInt
	for i := 0; i < len(array); i ++{
		if array[i] > 0 && array[i] < minPositive{
			minPositive = array[i]
		}else if array[i] < 0 && array[i] > maxNegative{
			maxNegative = array[i]	
		}
	}
	minIndex := -1
	maxIndex := -1
	for i := 0; i < len(array); i++{
		if array[i] == minPositive {
			minIndex = i
		  } 
	}
	for i := 0; i < len(array); i++{
		if array[i] == maxNegative {
			maxIndex = i
		  } 
	}
	if minIndex != -1{
		array = append(array[:minIndex], array[minIndex+1:]...)
	}
	if maxIndex != -1{
		array = append(array[:maxIndex], array[maxIndex+1:]...)
	}
	fmt.Println(array)

}
