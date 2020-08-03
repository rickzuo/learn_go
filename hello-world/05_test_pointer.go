package main

import "fmt"

func main() {
	num := 1

	p := &num

	pNum := *p
	fmt.Println(p,&p)
	fmt.Println(pNum)
	fmt.Printf("%T ",p)
	fmt.Printf("%T ",pNum)


	//var a *int // nil pointer
	var a = new(int)
	v := *a
	fmt.Printf("\n %d ",a)
	fmt.Printf("%T ",a)
	fmt.Printf("%d ",v)
	fmt.Println(*a)
}
