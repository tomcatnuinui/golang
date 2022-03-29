package main
import(
	"fmt"
)
func main() {
	var sr string
	fmt.Println("请输入要转的字母：")
	fmt.Scanln(&sr)
	switch sr {
		case "a":
			fmt.Println("A")
		case "b":
			fmt.Println("B")
		case "c":
			fmt.Println("C")
		case "d":
			fmt.Println("D")
		case "e":
			fmt.Println("E")
		default:
			fmt.Println("other")
	}

}