package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/edvaldo-domingos/go_banking/service"
	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	
	// customers := []Customer{
	// 	{"Samanta", "Cape Town", "1234"},
	// 	{"Jon", "New York", "4321"},
	// }

	customers, _ := ch.service.GetAllCustomers()
		
	if r.Header.Get("Content-Type") == "application/xml" {
		rw.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(rw).Encode(customers)

	}else{

		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(customers)
	}

}

func (ch *CustomerHandler) getCustomer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)

	if err != nil{
		writeResponse(rw, err.Code, err.AsMessage())
	}else{
		writeResponse(rw, http.StatusOK, customer)
	}
}

func writeResponse(rw http.ResponseWriter, code int, data interface{}){
	rw.Header().Add("Content-Type", "application.json")
	rw.WriteHeader(code)
	if err := json.NewEncoder(rw).Encode(data); err != nil {
		panic(err)
	}
}