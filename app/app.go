package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/edvaldo-domingos/go_banking/domain"
	"github.com/edvaldo-domingos/go_banking/logger"
	"github.com/edvaldo-domingos/go_banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)


func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}

	for _, k := range envProps {
		if os.Getenv(k) == "" {
			log.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}

}

func Start(){
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
	router := mux.NewRouter()

	//wiring
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	// accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	ch := CustomerHandler{service.NewCustomerService(customerRepositoryDb)}
	// ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	
	// defining routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logger.Info(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}


func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
