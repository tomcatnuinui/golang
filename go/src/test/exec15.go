package main
import(
	"fmt"
)

type MethodUtils struct {
	Cang int
	Kuan int
}

func (m MethodUtils) test() {
	for i :=0;i<=m.Cang;i++{
		fmt.Printf("*")
	} 
}

func main(){
	m := MethodUtils {
		Cang : 10,
		Kuan :8,
	}
	m.test()
}