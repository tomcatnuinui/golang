package main
import(
	"fmt"
)
type P struct {
	Name string
	Age int
}
var a P
func main() {
	fmt.Printf("%+v\n",a)
}