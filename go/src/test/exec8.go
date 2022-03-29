package main
import(
	"fmt"
)
func main(){
	a := make(map[int]map[string]string)
	a[1]=make(map[string]string)
	a[1]["name"]="skr"
	a[1]["addres"]="北京"
	a[1]["sex"]="女"

	fmt.Println(a[1]["sex"])
	
}