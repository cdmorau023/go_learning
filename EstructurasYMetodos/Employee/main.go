package main

import (
	"fmt"
	"time"
)

type Person struct {
	ID          int
	Name        string
	DateOfBirth time.Time
}
type Employee struct {
	ID       int
	Position string
	Person
}

func (e Employee) PrintEmployee() {
	fmt.Printf("ID: %d\nName: %s\nDate of Birth: %s\nPosition: %s\n", e.ID, e.Name, e.DateOfBirth, e.Position)
}

func main() {
	persona := Person{
		ID:          3,
		Name:        "Juan",
		DateOfBirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	empleado := Employee{
		ID:       1,
		Position: "Developer",
		Person:   persona,
	}

	empleado.PrintEmployee()
}
