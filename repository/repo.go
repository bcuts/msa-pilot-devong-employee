package repository

import (
	"fmt"
	"act-msa-pilot-devong-employee/model"
)

var currentId int

var employees model.Employees

// Give us some seed data
func init() {
	RepoCreateEmployee(model.Employee{FirstName: "Write presentation"})
	RepoCreateEmployee(model.Employee{FirstName: "Host meetup"})
}

func RepoFindEmployee(id int) model.Employee {
	for _, t := range employees {
		if t.ID == id {
			return t
		}
	}
	// return empty Employee if not found
	return model.Employee{}
}

func RepoFindEmployees() interface{} {
	return employees
}

func RepoCreateEmployee(employee model.Employee) model.Employee {
	currentId += 1
	employee.ID = currentId
	employees = append(employees, employee)
	return employee
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

func RepoUpdateEmployee(employee model.Employee) interface{} {
	for i, t := range employees {
		if t.ID == employee.ID {
			employees[i] = employee
			return employee
		}
	}
	return model.Employee{}
}
