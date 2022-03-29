package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p Person) speak() {
	fmt.Println(p.Name, "在说话")
	return
}

func (p Person) jisuan() {
	res := 0
	for i := 0; i <= 1000; i++ {
		res += i
	}
	fmt.Println("res =", res)
}

func (p Person) jisuan2(n int) {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	fmt.Println("res =", res)
}

func (p Person) getsum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	a := Person{
		Name: "tom",
	}
	a.speak()
	a.jisuan()
	a.jisuan2(100)
	fmt.Println(a.getsum(10, 20))
}
