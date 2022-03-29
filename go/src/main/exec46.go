package main

import (
	"fmt"
)
//二分查找
func efcz(a *[]int, left int, right int, find int) {
	if left > right {
		fmt.Println("找不到")
		return
	}
	meid := (left + right) / 2
	if (*a)[meid] > find {
		efcz(a, left, meid-1, find)
	} else if (*a)[meid] < find {
		efcz(a, meid+1, right, find)
	} else {
		fmt.Println("找到了,下标为", meid)
	}
}
func main() {
	var b []int = []int{23, 44, 55, 66, 77, 88, 99}
	efcz(&b, 0, 6, 99)
}
