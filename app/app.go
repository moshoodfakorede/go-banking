package app

import (
	"log"
	"net/http"

	"github.com/Fakorede/banking/domain"
	"github.com/Fakorede/banking/service"
	"github.com/gorilla/mux"
)

// Start app
func Start() {
	// custom multiplexer
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// start server
	log.Fatal(http.ListenAndServe("localhost:100", router))
}
