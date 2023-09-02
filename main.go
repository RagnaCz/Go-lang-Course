package main

import "fmt"

var numberInt, numberInt2 int = 1000, 2000
var msg string = "Hello"

//Array
var productname [4]string

//var price [4]float32

func main() {
	numberfloat := 25.4

	productname[0] = "mac"

	price := [4]float32{100, 200, 300, 400}

	fmt.Println("Hello Tanawat Kanchan")
	fmt.Println(numberInt)
	fmt.Println(numberfloat)
	fmt.Println(msg)

	fmt.Println(productname)
	fmt.Println(price)
}
