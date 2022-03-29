package main
//创建切片
import(
	"fmt"
)
func main() {
	var a = make([]string,4)
	a[0] = "fang"
	a[1] = "skr"
	a[2] = "sky"
	a[3] = "andy"
	fmt.Println(len(a))
	fmt.Println(a)
	var b = make([]string,2)
	copy(b,a[2:2])
	fmt.Println(len(b))
	fmt.Println(b)
	
}