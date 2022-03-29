package main
import(
	"fmt"
	"io/ioutil"
)
func main(){
	file,err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("%s",file)
}