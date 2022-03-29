package main
import(
	"fmt"
)

type Visitor struct {
	name string
	age int
}

func (v *Visitor) test(){
	
}

func main(){
	v := Visitor{
		name : "tom",
		age : 20,
	} 
	
	
}