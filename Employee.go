package main


type Employee struct {
	ID        string	`json:"id"`
	FirstName string	`json:"firstName"`
	LastName  string	`json:"lastName"`
	BranchId  string	`json:"branchId"`
}

type Employees []Employee


