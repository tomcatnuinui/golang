package main

import (
	"fmt"
	"math/rand"
	"time"
)
//二分查找函数
func efcz(a *[10]int, left int, right int, find int) {
	if left > right {
		fmt.Println("没找到")
		return
	}
	middle := (left + right) / 2
	if (*a)[middle] > find {
		efcz(a, left, middle-1, find)
	} else if (*a)[middle] < find {
		efcz(a, middle+1, right, find)
	} else {
		fmt.Printf("找到了，下标是%v\n", middle)
	}
}
func main() {
	var b [10]int
	//定义随机种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		b[i] = a
	}
	//生成1-100的随机数
	fmt.Println(b)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if b[i] < b[j] {
				b[i], b[j] = b[j], b[i]
			}
		}
	}
	fmt.Println(b)
	efcz(&b, 0, 10, 90)

}
