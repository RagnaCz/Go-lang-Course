package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string) float64 {
	fmt.Printf("%v", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must number only", prompt)
		panic(message)
	}
	return value
}

func add(value1 float64, value2 float64) float64 { return value2 + value1 }

func getOperator() string {
	fmt.Print("operator is ( + - * / ) : ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func main() {
	value1 := getInput(" value1 = ")
	value2 := getInput(" value2 = ")
	var result float64

	switch operator := getOperator(); operator {
	case "+":
		result = add(value1, value2)
	default:
		fmt.Println("error")
	}

	fmt.Println("result = ", result)

}
