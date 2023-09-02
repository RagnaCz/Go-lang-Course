package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {
	/*employee1 := employee{
		employeeID:   "101",
		employeeName: "Tanawat",
		phone:        "0611799925",
	}

	fmt.Println("employee1 = ", employee1)
	*/
	employeeList := [3]employee{}
	employeeList[0] = employee{
		employeeID:   "101",
		employeeName: "Hello",
		phone:        "222",
	}

	// can use slice

	fmt.Println(employeeList)
}
