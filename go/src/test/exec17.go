package main

import (
	"fmt"
)

type Calcuator struct {
	Num1 int
	Num2 int
}

func (c *Calcuator) print(key string) {
	for i := 1; i <= c.Num1; i++ {
		for j := 1; j <= c.Num2; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

func main() {
	c := Calcuator{
		Num1: 5,
		Num2: 6,
	}

	c.print("^")

}
