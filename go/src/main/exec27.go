package main

import (
	"fmt"
)

func main() {
	var month int
	fmt.Println("请输入现在的月份：")
	fmt.Scanln(&month)
	if month >= 4 && month <= 10 {
		//旺季
		var age int
		fmt.Println("请输入购票人年龄：")
		fmt.Scanln(&age)
		if age >= 18 && age <= 60 {
			fmt.Println("成人票价为60一位")
		} else if age < 18 {
			fmt.Println("儿童票价为30一位")
		} else {
			fmt.Println("老人票价为20一位")
		}
	} else {
		//淡季
		var age int
		fmt.Println("请输入购票人年龄：")
		fmt.Scanln(&age)
		if age >= 18 && age <= 60 {
			fmt.Println("成人票价为40一位")
		} else {
			fmt.Println("其他人票价为20一位")
		}
	}
}
