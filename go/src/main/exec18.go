package main

import (
	"fmt"
	"strconv"
)

type Movie struct {
	Name   string
	Rating float64
}

func (m *Movie) summary() string {
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
	return m.Name + "," + r
}

func main() {
	b := Movie{
		Name:   "2012",
		Rating: 3.2,
	}

	fmt.Println(b.summary())

}
