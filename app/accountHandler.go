package app

import (
	"encoding/json"
	"net/http"

	"github.com/xvbnm48/go-microservice-udemy/dto"
	"github.com/xvbnm48/go-microservice-udemy/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	} else {
		account, appErr := h.service.NewAccount(request)
		if appErr != nil {
			writeResponse(w, appErr.Code, appErr.AsMessage())
			return
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}

}
