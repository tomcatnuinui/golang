package main

import (
	"fmt"
	"time"
)

func slowfunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("sleeper() finished")
}

func main() {
	go slowfunc()
	fmt.Println("test test")
	time.Sleep(time.Second * 3)
}
