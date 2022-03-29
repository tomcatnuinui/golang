package main

import (
	"fmt"
)

func pwdf(a [8]float32) (float32, float32, float32, float32) {
	var max float32
	var min float32 = a[0]
	var total float32
	for i := 0; i < 8; i++ {
		total += a[i]
		if a[i] > max {
			max = a[i]
		} else if a[i] < min {
			min = a[i]
		}
	}
	return total, max, min, (total - max - min) / 6
}

func main() {
	var a [8]float32
	fmt.Println("请8位评委打分：")
	for i := 0; i < 8; i++ {
		fmt.Scanln(&a[i])
	}
	fmt.Println(a)
	b, c, d, e := pwdf(a)
	fmt.Printf("总分为%.1f,最高分为%.1f,最低分为%.1f,最后得分为%.2f\n", b, c, d, e)
	var max float32
	var min float32 
	var maxi int
	var mini int
	for i:=0;i<8;i++{
       if a[i] > e {
		(a[i] - e)=max
		maxi = i
	   } else {
		   (e - a[i])=max
		   maxi=i
	   }
	}

	fmt.Printf("最佳评委是%d,最差评委是%d\n",mini,maxi)
}
