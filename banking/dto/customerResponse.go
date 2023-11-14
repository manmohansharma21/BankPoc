package dto

type CustomerResponse struct {
	Id          string `json:"customer_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

/*
DTO layers helps in modularity and helps in prevening domain objects scattering all over the different layers = append(DTO layers helps in modularity and helps in prevening domain objects scattering all over the different layers = append(DTO layers helps in modularity and helps in prevening domain objects scattering all over the different layers = append(DTO layers helps in modularity and helps in prevening domain objects scattering all over the different layers.
As our domain should not be exposed to the outsider world, DTO layer will help here. Domain object and DTO hold different responsibilities where domain object is at service side layer and DTO is at user side layer.
DTO is for service layer(server side) object for data transformation whereas domain object is the user side object.
Domain has complete knowledge for constructing DTO.
*/
