package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct{
	Name string `json:"full_name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}


func greet(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello REST")
} 

func getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	
	customers := []Customer{
		{"Samanta", "Cape Town", "1234"},
		{"Jon", "New York", "4321"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		rw.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(rw).Encode(customers)

	}else{

		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(customers)
	}

}


func getCustomer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(rw, vars["customer_id"])
}

func createCustomer(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "POST request received")
}