package main

import (
	"log"
	"net/http"

	"github.com/hulliokaisar/logistic/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Endpoint untuk mengambil data kurir
	router.HandleFunc("/couriers", handlers.FetchCouriersHandler).Methods("GET")

	// Endpoint dengan parameter origin_name dan destination_name
	router.HandleFunc("/couriers/{origin_name}/{destination_name}", handlers.FilterCouriersHandler).Methods("GET")

	// Menyajikan file HTML
	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8000", router))
}
