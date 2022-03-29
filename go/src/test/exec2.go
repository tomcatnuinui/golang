package main

import (
	"fmt"
)

func main() {
	var a = [5]int{1, 9, 5, 7, 3}
	fmt.Println(a)

	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
			if a[i] > a[j] {
				a[i],a[j]=a[j],a[i]
			}
		}
	}
	fmt.Println(a)
}
