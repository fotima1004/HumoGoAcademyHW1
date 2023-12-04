package main

import "fmt"
func main () {
	var pol_count int = 0
	var otr_count int = 0
	var a int 
	fmt.Println("Enter the number for a:")
	fmt.Scan(&a)
	var b int 
	fmt.Println("Enter the number for b:")
	fmt.Scan(&b)
	var c int 
	fmt.Println("Enter the number for c:")
	fmt.Scan(&c)
	if a > 0{
		pol_count += 1
	}else if a < 0{
		otr_count += 1		
	}
	if b > 0{
		pol_count += 1
	}else if b < 0{
		otr_count += 1		
	}
	if c > 0{
		pol_count += 1
	}else if c < 0{
		otr_count += 1		
	}			
	fmt.Println("number of positive",pol_count)
	fmt.Println("number of negativewq",otr_count)				
}