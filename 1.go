package main

import (
	"fmt"

	"github.com/q1mi/hello"
)

const Pi = 3.14

func main() {
	const World = "世界"
	a := 1
	fmt.Println(a)
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	fmt.Print("Happy", Pi, "你好\n")
	const Truth = true
	fmt.Println("Go rules?", Truth)
	hello.SayHi()
	//b := "年后打工2"
	//fmt.Printf("b类型为：%T\n", b)
}
