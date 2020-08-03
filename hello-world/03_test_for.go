package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var j = 0
	var k int = 20
	fmt.Println("j = ", j)
	fmt.Println("k = ", k)
	for ; j < 5; j++ {
		fmt.Println(j)
	}

	var str = "乘风破浪"
	for index,v := range(str) {
		fmt.Println(index,v)
		fmt.Printf("%d,%c \n",index,v)
	}

	for {
		fmt.Println(666, time.Duration(2), time.Second, time.Duration(2)*time.Second)
		time.Sleep(time.Duration(2) * time.Second)
	}


}
