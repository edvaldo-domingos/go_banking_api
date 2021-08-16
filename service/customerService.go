package service

import (
	"github.com/ashishjuyal/banking-lib/errs"
	"github.com/edvaldo-domingos/go_banking/domain"
)

type CustomerService interface{
	GetAllCustomers(status string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}


type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// receiver function
func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError){
	if status == "active"{
		status ="1"
	}else if status == "inactive"{
		status = "0"
	}	else{
		status = ""
	}
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError){
	return s.repo.ById(id)
}

// business logic
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService{
	return DefaultCustomerService{repository}
}