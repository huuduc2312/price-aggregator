package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func singlePrice(w http.ResponseWriter, r *http.Request) {
	log.Println("-- Received request: singlePrice --")

	// Print the latest price of the specified asset
	w.Header().Set("Content-Type", "application/json")

	price, err := getBinanceBTCPrice()
	if err != nil {
		log.Println("get price failed:", err)
	}

	log.Println("--")

	json.NewEncoder(w).Encode(map[string]any{"result": price.open * 1e10})
}
