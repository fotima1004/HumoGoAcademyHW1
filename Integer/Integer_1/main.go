package main

import "fmt"

func main () {
	var L int
	fmt.Println("Enter the distance in centimeters:")
	fmt.Scan(&L)
	if L >= 100{
		var Metr = L / 100
		fmt.Println("Metr",Metr)
	}else {
		fmt.Println("Number do not meet requirements")	
	}
	
}