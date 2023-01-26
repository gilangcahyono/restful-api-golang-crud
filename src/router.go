package src

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", Home).Methods("GET", "OPTIONS")
	router.HandleFunc("/products", GetProduct).Methods("GET", "OPTIONS")
	router.HandleFunc("/products/{id}", GetProductById).Methods("GET", "OPTIONS")
	router.HandleFunc("/products", CreateProduct).Methods("POST", "OPTIONS")
	router.HandleFunc("/products/{id}", UpdateProduct).Methods("PUT", "OPTIONS")
	router.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE", "OPTIONS")

	return router
}
