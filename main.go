package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	// Define the endpoint to get the price
	http.HandleFunc("/price", priceAggregate)
	http.HandleFunc("/single-price", singlePrice)
	http.HandleFunc("/", wsHandler)

	// Start the HTTP server
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
