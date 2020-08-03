package main

import "fmt"

func main() {
	var i int = 10
	if i < 0{
		fmt.Println("i = ",i)
	}else if i % 10 == 0 {
		fmt.Println("i % 10 = 0,i = ",i)
	}else{
		fmt.Println("else i = ",i)
	}
}
