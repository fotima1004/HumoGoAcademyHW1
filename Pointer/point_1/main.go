package main
import "fmt"

func repLacePointers(a, b *int){
	*a,*b = *b,*a

}

func main()  {
	var a, b int
	fmt.Scan(&a,&b)
	repLacePointers(&a,&b)
	fmt.Println(a,b)
	
}
/*
import "fmt"

func pow2(a *int){
	*a *= *a

}

func main()  {
	var a int
	fmt.Scan(&a)
	pow2(&a)
	fmt.Println(a)
	
}*/