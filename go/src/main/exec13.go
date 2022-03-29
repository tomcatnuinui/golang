package main
import(
	"fmt"
	"time"
)
func main() {
	t := time.Now()
	fmt.Println(t.Format("20061001150000"))
}