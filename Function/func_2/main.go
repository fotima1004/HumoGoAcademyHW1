package main

import "fmt"

func vote(a,b,c bool) bool {
	
	return((a && b)|| (b && c)||(a && c))
}
func main()  {
	fmt.Println(vote(false,true,false))
}