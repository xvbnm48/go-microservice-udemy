package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xvbnm48/go-microservice-udemy/service"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")

}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{
			Name: "nabila gusmarlia", City: "indonesia", Zipcode: "112929",
		},
		{
			Name: "Sakura Endo", City: "Japan", Zipcode: "48484848",
		},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		// json.NewEncoder(w).Encode(customers)
		xml.NewEncoder(w).Encode(customers)

	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)

	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprint(w, err.Message)
		fmt.Println("error when get customer")
		return
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customer)
	}

}
