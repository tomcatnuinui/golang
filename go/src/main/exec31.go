package main
import(
	"fmt"
)
func main() {
	var cj float32
	fmt.Println("请输入学生成绩：")
	fmt.Scanln(&cj)
	switch {
		case cj > 100:
			fmt.Println("输入的成绩不能大于100分")
		case cj >= 60:
			fmt.Println("成绩合格")
		default:
			fmt.Println("成绩不合格")
	}
}