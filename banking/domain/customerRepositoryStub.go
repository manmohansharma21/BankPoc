package domain

// Adapter for the secondary port
// Defining a port like defining a protocol; any component following this protocol should be able to
// able to  connect this port
// Mock is used to test while writing behavior based tests, mock is different from stub which contains the hard-coded values.
// Go uses duck typing, i.e., if it looks like a duck or quacks like a duck, then it is a duck.
// So we just need to define the interface function with receiver as stub defined.
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// Helper function responsible for creating customers, to instantiate defaultCustomerService
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Manm", City: "Vrindavan", Zipcode: "281121", DateOfBirth: "1996-06-21", Status: "1"},
		{Id: "1002", Name: "Bhav", City: "Vrindavan", Zipcode: "281121", DateOfBirth: "1996-06-21", Status: "1"},
	}

	return CustomerRepositoryStub{customers: customers}
}

// Injected this stub implementation, later on, will inject real database adapter at the time of wiring our application.
