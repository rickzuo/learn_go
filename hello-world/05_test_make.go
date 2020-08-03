package main

import "fmt"

func main() {
	var a1 = make([]string,10)

	a1 = append(a1,"sdf")
	for index,v := range a1{
		fmt.Println(index,v)
	}
	fmt.Println("a1:",a1)
}
