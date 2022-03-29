package main

import (
	"fmt"
)

type Calcuator struct {
	Num1 float64
	Num2 float64
}

func (c *Calcuator) get(key string) float64 {
	res := 0.0
	switch key {
	case "-":
		res = c.Num1 - c.Num2
	case "+":
		res = c.Num1 + c.Num2
	case "*":
		res = c.Num1 * c.Num2
	case "/":
		res = c.Num1 / c.Num2
	default:
		fmt.Println("输入的运算符错误")
	}
	return res
}

func main() {
	a := Calcuator{
		Num1: 2.3,
		Num2: 2.2,
	}
	fmt.Println(fmt.Sprintf("%.2f", a.get("+")))
}
