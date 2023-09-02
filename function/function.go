package main

import "fmt"

func Hello() {
	fmt.Println("Hello Tanawat")
}

func plus(value1 int, value2 int) int {
	return value1 + value2
}

func plus3value(value1 int, value2 int, value3 int) int {
	return value1 + value2 + value3
}

func main() {
	Hello()
	result := plus(1, 2)
	fmt.Println("result = ", result)

	result2 := plus3value(5, 5, 10)
	fmt.Println("result 2 = ", result2)
}
