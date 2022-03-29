package main

import (
	"fmt"
	"bytes"
)

func main() {
	var b bytes.Buffer
	for i := 0; i < 500; i++ {
		b.WriteString("z")
	}
	fmt.Println(b.String())
}
