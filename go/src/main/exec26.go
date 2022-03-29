package main

import (
	"fmt"
)

func main() {
	var second float64

	fmt.Println("请输入比赛成绩(秒)：")
	fmt.Scanln(&second)
	if second <= 8 {
		fmt.Println("请输入性别：")
		var gender string
		fmt.Scanln(&gender)
		if gender == "男" {
			fmt.Println("进入男子组决赛")
		} else if gender == "女" {
			fmt.Println("进入女子组决赛")
		} else {
			fmt.Println("请输入正确的性别")
		}
	} else {
		fmt.Println("对不起，你没能进入决赛")
	}

}
