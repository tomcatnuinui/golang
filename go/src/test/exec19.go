package main
import(
	"fmt"
)

type MethodUtils struct {
	Num int
}

func (m *MethodUtils) print(n int) {
	for i:=1;i<=n;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%vx%v=%v",i,j,i*j)
		}
		fmt.Println()
	}
}

func main(){
	
}