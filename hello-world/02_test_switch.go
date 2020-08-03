package main

import "fmt"

func main() {
	num := 1

	switch num {
	case 1:
		fmt.Println("i am ", num)
		fallthrough // 不跳出，后面的无条件执行
	case 2,3,4:
		fmt.Println("i am ", num)
	default:
		fmt.Println("i am default ")
	}


	switch n := 10; n{
	case 10,1:
		fmt.Println("nice +")
	case 12:
		fmt.Println("nice")
	case 13:
		fmt.Println("nice -")
	default:
		fmt.Println("i am default ")
	}

	score := 89
	switch {
	case score > 90:
		fmt.Println("nice +")
	case score < 80:
		fmt.Println("nice")
	case score > 60:
		fmt.Println("nice -")
	default:
		fmt.Println("i am default ")
	}
}
