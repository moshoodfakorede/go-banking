package domain

// CustomerRepositoryStub struct
type CustomerRepositoryStub struct {
	customers []Customer
}

// FindAll return all customers
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// NewCustomerRepositoryStub function
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{ID: "001", Name: "John", City: "Wick", Zipcode: "110121", DateOfBirth: "2000-10-12", Status: "1"},
		{ID: "002", Name: "Brad", City: "Traversy", Zipcode: "110121", DateOfBirth: "2000-10-12", Status: "1"},
	}

	return CustomerRepositoryStub{customers: customers}
}
