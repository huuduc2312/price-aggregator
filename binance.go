package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"time"
)

type priceResult struct {
	open float64
	high float64
	low  float64
}

func getBinanceBTCPrice() (*priceResult, error) {
	// Define the Binance API URL for retrieving the latest price of a symbol
	apiURL := "https://api.binance.com/api/v3/klines?symbol=BTCUSDT&interval=1m&limit=1"

	t := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		// We use ABSURDLY large keys, and should probably not.
		TLSHandshakeTimeout: 60 * time.Second,
	}
	c := &http.Client{
		Transport: t,
	}

	// Make a GET request to the Binance API
	resp, err := c.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("Error making request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response: %w", err)
	}

	// Define a struct to unmarshal the JSON response
	var response [][]any

	// Unmarshal the JSON response into the defined struct
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling response: %w", err)
	}

	price := response[0]
	open, err := strconv.ParseFloat(fmt.Sprint(price[1]), 64)
	if err != nil {
		return nil, fmt.Errorf("parse open price failed: %w", err)
	}

	high, err := strconv.ParseFloat(fmt.Sprint(price[2]), 64)
	if err != nil {
		return nil, fmt.Errorf("parse high price failed: %w", err)
	}

	low, err := strconv.ParseFloat(fmt.Sprint(price[3]), 64)
	if err != nil {
		return nil, fmt.Errorf("parse low price failed: %w", err)
	}

	open = 68600
	high = 68600
	low = 68600

	return &priceResult{open: open, high: high, low: low}, nil
}
