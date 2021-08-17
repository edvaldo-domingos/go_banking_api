package service

import (
	"github.com/ashishjuyal/banking-lib/errs"
	"github.com/edvaldo-domingos/go_banking/domain"
	"github.com/edvaldo-domingos/go_banking/dto"
)

type CustomerService interface{
	GetAllCustomers(status string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
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

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError){
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response  := c.ToDto()
	
	return &response, nil
}

// business logic
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService{
	return DefaultCustomerService{repository}
}