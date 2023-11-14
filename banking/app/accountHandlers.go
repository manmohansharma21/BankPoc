package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manmohansharma21/bankpoc/banking/dto"
	"github.com/manmohansharma21/bankpoc/banking/service"
)

type AccountHandler struct {
	service service.AccountService //Handle for the service
}

func (h AccountHandler) NewAcount(w http.ResponseWriter, r *http.Request) {
	var request dto.NewAccountRequest

	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

// customers/2000/accounts/90720 --> example
func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	// Get the account_id and customer_id from the URL
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	// decode incoming request
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	}

	// build the request object
	request.AccountId = accountId
	request.CustomerId = customerId

	// make transaction
	account, appError := h.service.MakeTransaction(request)

	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, account)
	}

}
