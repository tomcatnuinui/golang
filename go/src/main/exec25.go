package main
import(
	"fmt"
)
func main() {
	var a float64
	var b float64
	var s string
	fmt.Println("请输入男方身高是多少：")
	fmt.Scanln(&a)
	fmt.Println("请输入男方家产有多少(单位万)：")
	fmt.Scanln(&b)
	fmt.Println("请输入男方长相帅不帅：")
	fmt.Scanln(&s)
	if a > 180.0 && b > 1000.0 && s == "帅" {
		fmt.Println("我一定嫁给他")
	} else if a > 180.0 || b > 1000.0 || s == "帅" {
		fmt.Println("嫁吧，比上不足，比下有余")
	} else {
		fmt.Println("坚诀不嫁")
	}
}