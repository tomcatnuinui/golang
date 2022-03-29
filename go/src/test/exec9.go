package main
import(
	"fmt"
)
func main(){
	var test []map[string]string
	test = make([]map[string]string,2)
	test[0]=make(map[string]string)
	test[0]["name"]="skr"
	test[0]["age"]="20"
	test[0]["address"]="chengdu"
	test[1]=make(map[string]string)
	test[1]["name"]="dog"
	test[1]["age"]="30"
	test[1]["address"]="chengdu"
	fmt.Println(test[1])
}