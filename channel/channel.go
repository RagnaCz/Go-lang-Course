package main

import (
	"fmt"
	"time"
)

//channel = thread

func f(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

//go and func = thread

func main() {
	go f("Hello")
	go f("world")
	time.Sleep(5 * time.Second)
}
