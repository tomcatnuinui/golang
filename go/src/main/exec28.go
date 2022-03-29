package main

import (
	"fmt"
)

func main() {
	var js string
	fmt.Println("请输入字母：")
	fmt.Scanln(&js)
	switch js {
		case "a","A":
			fmt.Println("星期一")
		// case 'b':
		// 	fmt.Println("星期二")
		// case 'c':
		// 	fmt.Println("星期三")
		// case 'd':
		// 	fmt.Println("星期四")
		// case 'e':
		// 	fmt.Println("星期五")
		// case 'f':
		// 	fmt.Println("星期六")
		// case 'g':
		// 	fmt.Println("星期日")
		default:
			fmt.Println("输入的值不符合规范")
	}
	fmt.Println(js)
}
