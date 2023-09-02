package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Java", "Python"}
	fmt.Println(courseName)
	courseName = append(courseName, "C", "C#", "HTML")
	fmt.Println(courseName)

	courseWeb := courseName[1:3]
	fmt.Println(courseWeb)
}
