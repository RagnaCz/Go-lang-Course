package main

import "fmt"

func main() {
	i := 0
	for i < 2 {
		fmt.Println("i = ", i)
		i++
	}
	var input string
	for {
		fmt.Scanf("%s", &input)
		fmt.Println("input = ", input)
		if input == "exit" {
			break
		}
	}
}
