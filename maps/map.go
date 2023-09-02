package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("Product = ", product)

	//Add value on map
	product["Macbook"] = 40000
	product["Iphone"] = 30000
	product["Ipad"] = 20000

	fmt.Println("Product = ", product)

	//delete with key
	delete(product, "Ipad")
	fmt.Println("Product = ", product)

	//update
	product["Iphone"] = 20000
	fmt.Println("Product = ", product)

	value1 := product["Iphone"]
	fmt.Println("value1 = ", value1)

	courseName := map[string]string{"101": "Java", "102": "C++"}
	fmt.Println(courseName)
}
