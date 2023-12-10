package main

import (
	"fmt"
	"time"
)

type Employee struct {
	id        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerId int
}

func main() {
	var robert Employee
	robert.Name = "Robert"
	robert.Position = "Software Engineer"
	fmt.Printf("%s's initial salary is $%d\n", robert.Name, robert.Salary) // Robert's initial salary is $0
	robert.Salary = 1000
	fmt.Printf("%s's current salary is $%d\n", robert.Name, robert.Salary) // Robert's current salary is $1000
	robert.Salary += 500
	fmt.Printf("%s's current salary after hike is $%d\n", robert.Name, robert.Salary) // Robert's current salary after hike is $1500
	// Access data through pointers
	position := &robert.Position
	fmt.Printf("%s's current position before appraisal is %s\n", robert.Name, robert.Position) // Robert's current position before appraisal is Software Engineer
	*position = "Senior " + *position
	fmt.Printf("%s's current position after appraisal is %s\n", robert.Name, robert.Position) // Robert's current position after appraisal is Senior Software Engineer
}
