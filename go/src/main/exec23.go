package main
import(
	"fmt"
	"strings"
)

func main() {
	s := "Oh I do like to be beside the seaside"
	n := strings.ToUpper(s)
	o := strings.Replace(s,"seaside","bar",1)
	fmt.Println(n)
	fmt.Println(o)
	fmt.Println(strings.Index(o,"the"))

}