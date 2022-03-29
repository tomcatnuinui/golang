package main

import (
	"fmt"
)

type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	var a Circle
	a.radius = 2
	res := a.area()
	fmt.Println(res)
}
