package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hulliokaisar/logistic/database"

	"github.com/gorilla/mux"
)

type Courier struct {
	ID                int     `json:"id"`
	LogisticName      string  `json:"logistic_name"`
	Amount            int     `json:"amount"`
	DestinationName   string  `json:"destination_name"`
	OriginName        string  `json:"origin_name"`
	Duration          string  `json:"duration"`
	Additional_column *string `json:"additional_column"`
}

func FetchCouriersHandler(w http.ResponseWriter, r *http.Request) {
	// Lakukan validasi JWT di sini

	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM couriers")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	couriers := []Courier{}
	for rows.Next() {
		courier := Courier{}
		err := rows.Scan(&courier.ID, &courier.LogisticName, &courier.Amount, &courier.DestinationName, &courier.OriginName, &courier.Duration, &courier.Additional_column)
		if err != nil {
			log.Fatal(err)
		}
		couriers = append(couriers, courier)
	}

	json.NewEncoder(w).Encode(couriers)
}

func FilterCouriersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	originName := vars["origin_name"]
	destinationName := vars["destination_name"]

	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM couriers WHERE origin_name = ? AND destination_name = ?", originName, destinationName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	couriers := []Courier{}
	for rows.Next() {
		courier := Courier{}
		err := rows.Scan(&courier.ID, &courier.LogisticName, &courier.Amount, &courier.DestinationName, &courier.OriginName, &courier.Duration, &courier.Additional_column)
		if err != nil {
			log.Fatal(err)
		}
		couriers = append(couriers, courier)
	}

	json.NewEncoder(w).Encode(couriers)
}
