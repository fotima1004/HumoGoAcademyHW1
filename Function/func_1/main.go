package main

import "fmt"

func min4(a,b,c,d int) int {
	min_of_four := min(a, min(b, min(c, d)))
	fmt.Println(" minimum:",min_of_four)
	return min_of_four
}
func pow(n int) int{
	return n * n
}
func main()  {
	var a,b,c,d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println("function min4 starting")
	n := min4(a,b,c,d)
	min4(a,b,c,d)
	fmt.Println("function min4 ending")
	fmt.Println("function pow starting")
	pow(n)

}