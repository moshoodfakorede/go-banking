package service

import "github.com/Fakorede/banking/domain"

// CustomerService interface
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

// DefaultCustomerService struct
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomer return customers
func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// NewCustomerService instantiate this service
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
