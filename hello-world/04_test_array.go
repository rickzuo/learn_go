package main

import "fmt"

func main() {
	var a1 [3]bool
	var a2 [3]int

	fmt.Println(a1, a2)

	a1 = [3]bool{true, true, true}

	a10 := [...]int{1, 23, 4, 5, 5}
	a11 := [...]int{0: 1, 4: 23}
	fmt.Println(a1, a10, a11)

	for i := 0; i < len(a10); i++ {
		fmt.Println(i)
	}

	for _,v := range a11 {
		fmt.Println(v)
	}
}
