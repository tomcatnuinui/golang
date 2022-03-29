package main
import(
	"fmt"
	"io/ioutil"
)
func main() {
	 file,err := ioutil.ReadFile("exec1.go")
	 if err != nil {
		 fmt.Println(err)
		 return
	 }
	 fmt.Println(string(file))
	 //fmt.Println("%s",file)
}