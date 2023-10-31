package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manmohansharma21/bankpoc/dto"
	"github.com/manmohansharma21/bankpoc/service"
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
