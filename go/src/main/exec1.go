package main

import (
	"fmt"
)

func Sayhello() string {
	return "Hello"
}
func main() {
	var a int = 200
	defer fmt.Println("我是defer语句，我是扫尾的")

	switch a {
	case 2:
		fmt.Println("a=2")
	case 100:
		fmt.Println("a=2")
	case 180:
		fmt.Println("a=2")
	default:
		fmt.Println("a不知道是几")
	}

	fmt.Println("2")
	fmt.Println("1")
	fmt.Println(Sayhello())

}
