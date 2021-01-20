package domain

import "github.com/Fakorede/banking/errs"

// Customer struct
type Customer struct {
	ID          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// CustomerRepository repository
type CustomerRepository interface {
	// 1 == active, 0 == inactive
	FindAll(status string) ([]Customer, *errs.APIErrors)
	FindByID(string) (*Customer, *errs.APIErrors)
}
