package domain

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
	FindAll() ([]Customer, error)
}
