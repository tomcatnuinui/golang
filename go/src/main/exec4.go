package main
import(
	"fmt"
)
func main() {
	//切片
	var chess = make([]string,2)
	chess[0] = "Hello"
	chess[1] = "World"
	chess = append(chess,"test","dfsf","sfsf","dsfs")
	fmt.Println(chess)
}