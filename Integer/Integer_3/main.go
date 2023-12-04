package main

import "fmt"

func main () {
	var bt int
	fmt.Println("Enter the file size in bytes:")
	fmt.Scan(&bt)
	if bt >= 1024{
		var kb = bt / 1024
		fmt.Println("This file takes up ",kb," full kilobytes.")	
	}else {
		fmt.Println("Number do not meet requirements")
	}
}	