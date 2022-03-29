package main
//创建映射
import(
	"fmt"
)
func main() {
	var a = make(map[string]string)
	a["p"] = "段落"
	a["img"] = "图像"
	a["h1"] = "一级标题"
	a["h2"] = "二级标题"
	fmt.Println(len(a))
	fmt.Println(a)
	fmt.Println(a["p"])
}