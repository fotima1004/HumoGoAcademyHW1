package main

import "fmt"

func hhmmss_to_second(h1,m1,s1,h2,m2,s2 int) (h,m,s int){
	t1 := (h1 * 3600 + m1 * 60 +s1);
	t2 := (h2 * 3600 + m2 * 60 +s2);
	t3 := t2 - t1;
	h = t3 / 3600
	m = (t3 % 3600)/60
	s = t3 % 3600 % 60
	return h,m,s
}

func main()  {
	for i := 0; i < 2; i++{
		var h1,m1,s1,h2,m2,s2 int
		fmt.Println("Enter start time")
		fmt.Scan(&h1,&m1,&s1)
		fmt.Println("Enter the finish time")
		fmt.Scan(&h2,&m2,&s2)
		fmt.Println(hhmmss_to_second(h1,m1,s1,h2,m2,s2))
	}	
}