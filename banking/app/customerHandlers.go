package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manmohansharma21/bankpoc/banking/service"
)

// DTO: data Transfer Object
type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

/*
Can hit any endpoint either using 'curl' command on terminal or localhost over the browser or on Postman app.
routes==endpoints==patterns==apis on high level.
manmohansharma@Manmohans-MacBook-Air ~ % curl http://localhost:3055/customers
[{"full_name":"Manmohan","city":"Vrindavan","zipcode":"281121"},{"full_name":"Bhavesh","city":"Vrindavan","zipcode":"281121"}]
*/
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "Manmohan", City: "Vrindavan", Zipcode: "281121"},
	// 	{Name: "Bhavesh", City: "Vrindavan", Zipcode: "281121"},
	// }

	customers, _ := ch.service.GetAllCustomers("")

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	// Segment names are used to create a map of route variables, and can be retrieved using mux.Vars(r) where we need to parse the request inside it,
	// so this function returns map of all the segment names.
	vars := mux.Vars(r) //vars contains all the segments from the URL.
	//fmt.Fprint(w, vars["customer_id"])
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		//w.WriteHeader(http.StatusNotFound)
		//fmt.Fprintf(w, err.Error())

		// Order to follow: First header then error/http code then encode the JSON object else
		// else the content type is not applied to the response.
		//Error should also be in JSON if the response is encoded in JSON.
		writeResponse(w, err.Code, err.AsMessage()) // To use DRY principle, i.e., DO NOT REPEAT YOURSELF.
	} else {
		writeResponse(w, http.StatusOK, customer) // To use DRY principle, i.e., DO NOT REPEAT YOURSELF.
		// w.Header().Add("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode(customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, body interface{}) { //any is an alias for interface{} and is equivalent to interface{} in all ways.
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(body); err != nil {
		panic(err)
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post resuest received for customers")
}

func (ch *CustomerHandlers) getAllCustomersJSON(w http.ResponseWriter, r *http.Request) {
	// We should start with domain(means server side) for any changes in the code, and then service (client or primary port) side.

	// To retrieve Query Parameter from the request URL, i.e., /customers?status=active or /customer?status=inactive
	status := r.URL.Query().Get("status") // example: localhost:3077/customers?status=active
	//status := ""

	customers, err := ch.service.GetAllCustomers(status)

	if err != nil {
		//fmt.Fprint(w, err)
		writeResponse(w, err.Code, err.AsMessage()) // To use DRY principle, i.e., DO NOT REPEAT YOURSELF.
	} else {
		writeResponse(w, http.StatusOK, customers) // To use DRY principle, i.e., DO NOT REPEAT YOURSELF.
	}
}
