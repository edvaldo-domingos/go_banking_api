package service

import "github.com/edvaldo-domingos/go_banking/domain"

type CustomerService interface{
	GetAllCustomers() ([]domain.Customer, error)
}


type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// receiver function
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error){
	return s.repo.FindAll()
}

// business logic
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService{
	return DefaultCustomerService{repository}
}