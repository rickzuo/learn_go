package main

import "fmt"

func main() {
	var num int

	num = 10
	fmt.Println("num = ",num)

	var count int = 20

	count = 21
	fmt.Println("count:",count)

	flag := 31 // 推导赋值

	fmt.Printf("flag type is %T, %c",flag,flag)

	const c1 int = 10
	// error c1 = 20
	fmt.Println("c1 = ",c1)

	const c2 = 20 // no use :=
	fmt.Println("c2 = ",c2)

	var str string = "test"

	fmt.Println("str = ",len(str))
}
