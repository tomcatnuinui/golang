package main
import(
	"fmt"
)
func main() {
	//数组
	var chess [3]string
	chess[0] = "hello"
	chess[1] = "world"
	chess[2] = "test"
	fmt.Println(chess[0])
	fmt.Println(chess[1])
	fmt.Println(chess)
}
