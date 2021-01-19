package service

import (
	"github.com/Fakorede/banking/domain"
	"github.com/Fakorede/banking/errs"
)

// CustomerService interface
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, *errs.APIErrors)
}

// DefaultCustomerService struct
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomers function
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// GetCustomer return customers by id
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.APIErrors) {
	return s.repo.FindByID(id)
}

// NewCustomerService instantiate this service
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
