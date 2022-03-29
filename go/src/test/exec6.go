package main

import (
	"fmt"
)

func test(a []int) (int, int, int) {
	var total int
	var max int
	var min int
	for i := 0; i < len(a); i++ {
		total += a[i]
	}
	for i := 0; i < len(a); i++ {
		if a[i] > total/len(a) {
			max += 1
		} else if a[i] < total/len(a) {
			min += 1
		}
	}
	fmt.Println(total)

	return total / len(a), max, min
}

func main() {
	e := []int{66, 44, 54, 34, 55, 12, 33, 43}
	a, b, c := test(e)
	fmt.Println(a, b, c)
}
