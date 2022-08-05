package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xvbnm48/go-microservice-udemy/dto"
	"github.com/xvbnm48/go-microservice-udemy/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cusomterId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	} else {
		request.CustomerId = cusomterId
		account, appErr := h.service.NewAccount(request)
		if appErr != nil {
			writeResponse(w, appErr.Code, appErr.AsMessage())
			return
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}

}

func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	// get account id from url
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	// decode request body into transaction dto
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	} else {
		// build the request object
		request.AccountId = accountId
		request.AccountId = customerId

		// make transaction
		account, appError := h.service.MakeTRansaction(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
			return
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}
}
