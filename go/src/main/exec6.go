package main
import(
	"fmt"
)
func main() {
	var a = make([]string,3)
	a[0] = "you"
	a[1] = "I"
	a[2] = "he"
	var b = make([]string,2)
	// copy(b,a)
	// fmt.Println(len(b))
	// fmt.Println(b)
	copy(b,a[0:2])
	b = append(b,a[2])
	fmt.Println(len(b))
	fmt.Println(b)
}