package model

type Employee struct {
	ID        int    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	BranchId  string    `json:"branchId"`
}

type Employees []Employee