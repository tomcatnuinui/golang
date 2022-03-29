package main

import (
	"fmt"
)

func main() {
	var x interface{}
	var y = "a"
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型～ :%T", i)
	case int:
		fmt.Println("x 是int类型")
	case float64:
		fmt.Println("x 是float64类型")
	case func(int) float64:
		fmt.Println("x 是func(int)类型")
	case bool,string:
		fmt.Println("x 是bool或string型")
	default:
		fmt.Println("x 是未知类型")
	}
}
