package main

import "fmt"

func main() {

	var a1 []string
	var a2 []int

	fmt.Println(a1, a2)
	fmt.Println(a1 == nil)

	a1 = []string{"北京", "长沙", "上海"}
	a2 = []int{1, 2, 3, 54}
	fmt.Printf("a1 ,len = %d,cap = %d \n", len(a1), cap(a1))
	fmt.Printf("a2 ,len = %d,cap = %d \n", len(a2), cap(a2))

	var arr = [...]int{1, 23, 4, 5, 5}

	arrSlice := arr[0:3]
	fmt.Println("arr_slice = ", arrSlice , "arr = ",arr)
}
