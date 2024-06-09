package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	go getWsPrice(conn)
}

func getWsPrice(ws *websocket.Conn) {
	defer ws.Close()

	for {
		price, err := getBinanceBTCPrice()
		if err != nil {
			log.Println("get price failed:", err)
			continue
		}

		jsonData, err := json.Marshal([]float64{0, price.open})
		if err != nil {
			log.Println("Error marshalling JSON:", err)
			return
		}

		if err := ws.WriteMessage(websocket.TextMessage, jsonData); err != nil {
			log.Println("write ws data failed:", err)
			return
		}

		time.Sleep(1 * time.Second)
	}
}
