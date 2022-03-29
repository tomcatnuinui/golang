package main

import (
	"fmt"
)

type MethodUtils struct {
}

func (mu *MethodUtils) print() {
	for i := 1; i <= 10; i++ {
		for i := 1; i <= 8; i++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func (mu2 *MethodUtils) print2(m int, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func (mu3 *MethodUtils) area(m float64, n float64) float64 {
	return m * n
}

func (mu4 *MethodUtils) judge(m int) {
	if m%2 == 0 {
		fmt.Printf("%v是偶数\n", m)
	} else {
		fmt.Printf("%v是奇数\n", m)
	}
}

func main() {
	var m MethodUtils
	m.print()
	fmt.Println()
	m.print2(20, 8)

	res := m.area(5.0, 10.2)
	fmt.Println()
	fmt.Println("面积=", res)
	m.judge(10)

}
