package main

import "fmt"

func sayHi(str string)(string){

	msg := "Hi " + str
	fmt.Println(msg)
	return msg

}
type binOp func(int, int) int


func main() {
	msg := sayHi("rick")
	fmt.Println("msg " + msg)

}
