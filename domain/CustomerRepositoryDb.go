package domain

import (
	"database/sql"
	"time"

	"github.com/Fakorede/banking/errs"
	"github.com/Fakorede/banking/logger"
	_ "github.com/go-sql-driver/mysql"
)

// CustomerRepositoryDb struct
type CustomerRepositoryDb struct {
	client *sql.DB
}

// FindAll receiver function
func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.APIErrors) {
	// query db
	var rows *sql.Rows
	var err error

	if status == "" {
		findAllSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
		rows, err = d.client.Query(findAllSQL)
	} else {
		findAllSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE status = ?"
		rows, err = d.client.Query(findAllSQL, status)
	}

	if err != nil {
		logger.Error("Error while querying the customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			logger.Error("Error while scanning customers " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

		customers = append(customers, c)
	}

	return customers, nil

}

// FindByID receiver func
func (d CustomerRepositoryDb) FindByID(id string) (*Customer, *errs.APIErrors) {
	// query db
	customerSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"

	row := d.client.QueryRow(customerSQL, id)
	var c Customer
	err := row.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while querying the customers table " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

// NewCustomerRepositoryDb func
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
