package service

import (
	"github.com/ashishjuyal/banking-lib/errs"
	"github.com/edvaldo-domingos/go_banking/domain"
)

type CustomerService interface{
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}


type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// receiver function
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error){
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError){
	return s.repo.ById(id)
}

// business logic
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService{
	return DefaultCustomerService{repository}
}