package service

import (
	"github.com/Fakorede/banking/domain"
	"github.com/Fakorede/banking/errs"
)

// CustomerService interface
type CustomerService interface {
	GetAllCustomers(string) ([]domain.Customer, *errs.APIErrors)
	GetCustomer(string) (*domain.Customer, *errs.APIErrors)
}

// DefaultCustomerService struct
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomers return customers
func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.APIErrors) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	return s.repo.FindAll(status)
}

// GetCustomer return customers by id
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.APIErrors) {
	return s.repo.FindByID(id)
}

// NewCustomerService instantiate this service
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
