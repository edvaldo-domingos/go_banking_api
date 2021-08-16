package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/edvaldo-domingos/go_banking/service"
	"github.com/gorilla/mux"
)

type Customer struct{
	Name string `json:"full_name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

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
		rw.WriteHeader(err.Code)
		fmt.Fprintf(rw, err.Message)
	}else{
		rw.Header().Add("Content-Type", "application.json")
		json.NewEncoder(rw).Encode(customer)
	}

}