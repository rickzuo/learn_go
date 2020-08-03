package main

import "fmt"

func main() {

	var b bool = true

	var a string = "Runoob"
	fmt.Println(a)

	var c, d int = 1, 2
	if b {
		fmt.Println(c, d, b)
	}

	var i int
	var s string
	fmt.Printf("%v %v %v %q\n", i, b, s)


	f := "Runoob" // var f string = "Runoob"
	fmt.Println(f)

	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
