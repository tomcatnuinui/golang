package main
import(
	"fmt"
)
func cz(a [5]int)(int,int,int,int){
	var max int
	var maxindex int
    var min int
    var minindex int
	for i:=0;i<5;i++{
		min = a[0]
		if a[i] > max {
			max = a[i]
			maxindex = i
		} else if a[i] < min {
			min = a[i]
			minindex = i
		}
	}
	return max,maxindex,min,minindex
}

func main(){
	b := [5]int{3,4,2,190,1}
	max,maxindex,min,minindex := cz(b)
	fmt.Printf("最大的值是%d下标是%d,最小的值是%d下标是%d\n",max,maxindex,min,minindex)
}