package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	data, _ := json.Marshal(&employee{101, "Tanawat Kanchan", "0611799925", "golfangryb@gmail.com"})
	fmt.Println(data)
	fmt.Println(string(data))
}
