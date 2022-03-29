package main
import(
	"fmt"
)
type Pe struct {
	Name string
	Age int
}

func main() {
	// p := Pe {
	// 	Name: "skr",
	// 	Age: 30,
	// }
	var a Pe
	a.Name = "skr"
	a.Age = 30
	
	fmt.Println(a.Name)
}

