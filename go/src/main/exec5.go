package main
import(
	"fmt"
)
func main() {
	//创建切片
	var cheeses = make([]string ,3)
	cheeses[0] = "hello"
	cheeses[1] = "world"
	cheeses[2] = "chengdu"
	fmt.Println(len(cheeses))
	//追加数据到切片
	cheeses = append(cheeses,"11","22","33")
	fmt.Println(cheeses)
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
	//cheeses = append(cheeses[:3],cheeses[3+1:]...)
	//cheeses = append(cheeses[:4])
	//删除下标为2的数据
	cheeses = append(cheeses[:2],cheeses[2+1:]...)

	fmt.Println(cheeses)
}