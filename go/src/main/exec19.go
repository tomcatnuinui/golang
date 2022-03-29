package main

import (
	"fmt"
)

type Mj struct {
	chang float64
	kuan  float64
}

func (M *Mj) area() float64 {
	return 0.5 * (M.chang * M.kuan)
}
func (M *Mj) changechang(b float64) {
	M.chang = b
	return
}

func main() {
	a := Mj{chang: 6, kuan: 5}
	a.changechang(10)
	fmt.Println(a.area())
}
