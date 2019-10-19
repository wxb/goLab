package main

import (
	"fmt"
	"time"
)

// Employee 表述员工
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	var dilbert Employee
	fmt.Println(dilbert, &dilbert, &dilbert.Name, &dilbert.ID)

	position := &dilbert.Position
	fmt.Println(position, *position)
	*position += "Senior "
	fmt.Println(position, *position)

	// employeeOfTheMonth := &Employee{}
	// var employeeOfTheMonth *Employee = &dilbert
	var employeeOfTheMonth = &dilbert
	fmt.Println(employeeOfTheMonth, *employeeOfTheMonth)

	employeeOfTheMonth.Position += " (proactive team player)"
	(*employeeOfTheMonth).Salary = 29
	fmt.Println(*employeeOfTheMonth, dilbert)

}
