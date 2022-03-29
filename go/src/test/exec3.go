package main
import(
	"fmt"
)
func main(){
	var a =[10]string{"bb","cc","aa","dd","AA","AAAAA","aaa","ddf","AA","fdfd"}
	for i:=0;i<10;i++{
		if a[i] == "AA" {
			fmt.Println(i)
		}
	}

}