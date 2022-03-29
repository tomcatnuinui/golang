package main

import (
	"fmt"
)

type Game struct {
	Name string
	Date string
	Plat Plat
}
type Plat struct {
	XT   string
	CPU  string
	MEM  string
	Data string
}

func main() {
	a := Game{
		Name: "CS",
		Date: "2001",
		Plat: Plat{
			XT:   "windows",
			CPU:  "kengpeng",
			MEM:  "16G",
			Data: "500G",
		},
	}
	fmt.Printf("%+v\n",a)
	fmt.Println(a)
}
