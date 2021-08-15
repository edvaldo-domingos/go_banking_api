package app

import (
	"log"
	"net/http"

	"github.com/edvaldo-domingos/go_banking/domain"
	"github.com/edvaldo-domingos/go_banking/service"
	"github.com/gorilla/mux"
)


func Start(){
	router := mux.NewRouter()

	//wiring
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	
	// defining routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8080", router)) 

}

