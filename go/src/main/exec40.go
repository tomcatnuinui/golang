package main
import(
	"fmt"
)
func SumSub(n1,n2 int) (sum,sub int){
	sum = n1 + n2
	sub = n1 - n2
	return
}
func main(){
	fmt.Println(SumSub(30,10))
}