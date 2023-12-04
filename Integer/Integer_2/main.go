package main

import "fmt"

func main () {
	var M int
	fmt.Println("Enter mass in in kilograms:")
	fmt.Scan(&M)
	if M >= 1000{
		var Ton = M / 1000
		fmt.Println(Ton,"Tonna")
			
	}else {
		fmt.Println("Number do not meet requirements")	
	}
}	