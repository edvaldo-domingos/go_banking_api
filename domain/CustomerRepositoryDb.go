package domain

import (
	"database/sql"
	"time"

	"github.com/ashishjuyal/banking-lib/errs"
	"github.com/edvaldo-domingos/go_banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


type CustomerRepositoryDb struct {
	client *sqlx.DB
}


func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError){
	var err error
	customers := make([]Customer, 0)

	if status == ""{
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
	}else{
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)

	}

	if err != nil {
		logger.Error("Error while quering customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}


func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError){
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id =  ?"
	var c Customer

	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows{
			return nil, errs.NewNotFoundError("Customer not found")
		}else{
			logger.Error("Error while Scanning customer" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

	}

	return &c, nil

}

func NewCustomerRepositoryDb() CustomerRepositoryDb{
	client, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}

}