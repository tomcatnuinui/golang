package main
import(
	"fmt"
)
func main(){
	var a [4][6]int
	a[1][2] = 1
	a[2][3] = 2
	a[2][4] = 3

	for i := 0;i <len(a);i++{
		for j:=0;j<6;j++{
			fmt.Print(a[i][j]," ")
		}
		fmt.Println()
	}
	
}