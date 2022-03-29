package main
import(
	"fmt"
	"errors"
)
func main() {
	err := errors.New("This is test wrong")
	if err != nil {
		fmt.Println(err)
		//return
	}
	fmt.Println(123456)
}