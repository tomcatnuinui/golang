package main
import(
	"fmt"
)
func fbn(n int) ([]uint64){
	var fbnslise = make([]uint64,2)
	fbnslise[0]=1
	fbnslise[1]=1
	for i:=2;i<n;i++{
		fbnslise = append(fbnslise,fbnslise[i-1] + fbnslise[i-2])
	}
	return fbnslise

}
func main(){
	a := fbn(10)
	fmt.Println(a)
}