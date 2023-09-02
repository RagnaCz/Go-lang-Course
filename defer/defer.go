package main

import "fmt"

func add(val1, val2 float64) {
	result := val1 + val2
	fmt.Println("result :", result)
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}
}

func deferloop() {
	for j := 0; j < 10; j++ {
		defer fmt.Println("j : ", j)
	}
}

func main() {
	fmt.Println("Welcome to Calculator")
	defer fmt.Println("End") // move to the end of process
	add(20, 10)
	defer add(15, 13) //second
	defer add(12, 12) //first
	//fmt.Println("End")
	loop()
	deferloop()
}
