package main
import(
	"fmt"
)
func main(){
a:=[]string{"skr","sky","dog"}
fmt.Println(len(a))
a=append(a,"dddd")
fmt.Println(len(a))
fmt.Println(a)
}