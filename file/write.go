package main

import "os"

func main() {
	data1 := []byte("Hello\nWorld")
	err := os.WriteFile("data.txt", data1, 0664)

	if err != nil {
		panic(err)
	}

	f, ferr := os.Create("employeeName")
	if err != nil {
		panic(ferr)
	}
	defer f.Close()

	data2 := []byte("Tanawat\nKanchan")
	os.WriteFile("employeeName.txt", data2, 0644)
}
