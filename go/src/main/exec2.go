package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		for b := 1; b < 10; b++ {
			c := i * b
			if b <= i {
				fmt.Printf("%d X %d = %d\t", b, i, c)
			}
			if b == 9 {
				fmt.Println("\n")
			}
		}

	}
}
