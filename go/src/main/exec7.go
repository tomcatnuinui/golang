package main
import(
	"fmt"
)
func main() {
	var a = make(map[string]int)
	a["skr"] = 29
	a["fire"] = 30
	a["sky"] = 31
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(a["skr"])
}