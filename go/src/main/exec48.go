package main
import(
	"model"
	"fmt"
)

func main(){
	p := model.Newperson("tom")
	p.SetAge(20)
	a:=p.GetAge()
	fmt.Println(a)
}