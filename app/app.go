package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func Start(){
	router := mux.NewRouter()
	
	// defining routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id}", getCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8080", router)) 

}

