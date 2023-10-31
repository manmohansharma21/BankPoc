package service

import (
	"github.com/manmohansharma21/bankpoc/domain"
	"github.com/manmohansharma21/bankpoc/dto"
	"github.com/manmohansharma21/bankpoc/errs"
)

// This is service layer.

// Primary Port, i.e., facing to users
// All the ports are actually interfaces in Hexagonal architecture.
type CustomerService interface { //Name should reflect the intent.
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct { //For Implementation
	repo domain.CustomerRepository //DEPENDENCY
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	c, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	var response []dto.CustomerResponse

	for i := range c {
		response = append(response, c[i].ToDTO())
	}

	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindById(id) //Connected primary port with secondary port bymaking this call at secondary port
	if err != nil {
		return nil, err
	}

	//Transformations, i,e., fill all fields. //Domain has complete knowledge for constructing DTO.
	response := c.ToDTO() // Adding transfromations faciltated by the domain object. Domain will give DTO representation.
	return &response, nil
}

// Helper function to instantiate defaultCustomerService struct
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService { //Taking dependency as parameter
	return DefaultCustomerService{repo: repository}
}
