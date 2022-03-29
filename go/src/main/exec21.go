package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "Hello" + "world"
	a := 1233
	b := strconv.Itoa(a)
	fmt.Println(s+b)
}
