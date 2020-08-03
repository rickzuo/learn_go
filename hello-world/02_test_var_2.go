package main

import "fmt"

func test()(a, b, c int) {

	return 1,2,3
}
func main() {

	a := 10
	b := 20
	a, b = b, a
	fmt.Printf("a = %d , b = %d \n", a, b)

	_, d, _ := test() //匿名变量
	fmt.Printf("d = %d", d)

}
