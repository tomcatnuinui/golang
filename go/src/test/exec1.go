package main

//随机生成10个整数（1_100的范围）保存到数，并倒序打印以及平均值、求最大值和最大值的下标、并查找里面是否有55
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//定义一个数组，将随机出来的10个整数放进数组内
	var b [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		b[i] = a
	}
	fmt.Println(b)
	//用冒泡排序法排序并倒序打印
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if b[i] > b[j] {
				b[i], b[j] = b[j], b[i]
			}
		}
	}
	fmt.Println(b)

	sum := 0
	max := 0
	maxindex := 0
	//求数组整数的和
	for _, v := range b {
		sum += v
	}
	//求数组中最大值和最大值下标
	for i := 0; i < 10; i++ {
		if b[i] > max {
			max = b[i]
			maxindex = i
		}
	}
	//查找数组中是否有55这个整数
	for i := 0; i < 10; i++ {
		if b[i] == 55 {
			fmt.Println("数值中包含有55这个整数")
			break
		} else if i == 9 {
			fmt.Println("数值中未找到55这个整数")
		}
	}
	fmt.Printf("平均值是%v,最大值是%v,最大值的下标是%v\n", sum/10, max, maxindex)
}
