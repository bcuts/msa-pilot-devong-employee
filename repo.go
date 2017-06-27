package main

import (
	"fmt"
)

var currentId int

var employees Employees

// Give us some seed data
func init() {
	RepoCreateEmployee(Employee{FirstName: "Write presentation"})
	RepoCreateEmployee(Employee{FirstName: "Host meetup"})
}

func RepoFindEmployee(id int) Employee {
	for _, t := range employees {
		if t.ID == id {
			return t
		}
	}
	// return empty Employee if not found
	return Employee{}
}

func RepoFindEmployees() interface{} {
	return employees
}

func RepoCreateEmployee(t Employee) Employee {
	currentId += 1
	t.ID = currentId
	employees = append(employees, t)
	return t
}

func RepoDestroyEmployee(id int) error {
	for i, t := range employees {
		if t.ID == id {
			employees = append(employees[:i], employees[i + 1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Employee with id of %d to delete", id)
}

func RepoUpdateEmployee(employee Employee) interface{} {
	for i, t := range employees {
		if t.ID == employee.ID {
			employees[i] = employee
			return employee
		}
	}
	return Employee{}
}
