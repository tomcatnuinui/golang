package main

import (
	"fmt"
)

func main() {
	n, s := "skr", "bigdog"
	err := fmt.Errorf("this is %v and %v", n, s)
	if err != nil {
		fmt.Println(err)
	}
	panic("oh my god")
	a := 17 % 2
	fmt.Println(a)
}
