package domain

import (
	"github.com/manmohansharma21/bankpoc/dto"
	"github.com/manmohansharma21/bankpoc/errs"
)

// Secondary Port, i.e., facing to server infrastructure
// Business object definition
type Customer struct {
	Id          string `db:"customer_id"` //Property needs to match with table column names to be used in sqlx function calls. Hints to SQLX marshaller.
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusAsText() string {
	// Attach this knowledge to the domain object.
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}

	return statusAsText
}

// Domain logic
func (c Customer) ToDTO() dto.CustomerResponse {

	//Transformations, i,e., fill all fields.
	return dto.CustomerResponse{ //Domain has complete knowledge for constructing DTO.
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(), //Attached to domain object, status and their meanings. Keeping the logic at one place.
	}
}

// Define Repository Port in domain, this is secondary port for our hexagonal architecture
// This is our port which acts like a protocol which need to be abided to be used by an adapter
type CustomerRepository interface { //right intent names
	// Status==1 || Status==0 || Status =""
	FindAll(string) ([]Customer, *errs.AppError) //The object Implementing this becomes an adapter.
	FindById(string) (*Customer, *errs.AppError)
}

/*
	We should start with domain for any changes in the code.
	Domain side == Server side
*/
/*

This is all to facilitate Business logic decoupled to external work, agnostic to dependencies
This is how we can keep developing code without waiting for the database to be ready for use.
*/
