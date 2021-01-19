package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/Fakorede/banking/service"
)

// Customer struct
type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip"`
}

// CustomerHandlers struct
type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "John", City: "Wick", Zipcode: "110121"},
	// 	{Name: "Brad", City: "Traversy", Zipcode: "110121"},
	// }
	customers, _ := ch.service.GetAllCustomer()

	// content header
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
