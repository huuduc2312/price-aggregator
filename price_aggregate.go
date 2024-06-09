package main

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"
)

func priceAggregate(w http.ResponseWriter, r *http.Request) {
	log.Println("-- Received request: priceAggregate --")

	// Print the latest price of the specified asset
	w.Header().Set("Content-Type", "application/json")

	price, err := getBinanceBTCPrice()
	if err != nil {
		log.Println("get price failed:", err)
	}

	log.Println("open:", price.open)
	log.Println("high:", price.high)
	log.Println("low:", price.low)
	log.Println("--")

	// packed := packUint64ToUint256(open, high, low, 0)
	packed := pack64To256(uint64(price.open*1e10), uint64(price.high*1e10), uint64(price.low*1e10), 0)

	json.NewEncoder(w).Encode(map[string]string{"result": packed.String()})
}

func pack64To256(a, b, c, d uint64) *big.Int {
	packed := new(big.Int)
	packed.SetUint64(d)
	packed.Lsh(packed, 64)
	packed.Or(packed, big.NewInt(int64(c)))
	packed.Lsh(packed, 64)
	packed.Or(packed, big.NewInt(int64(b)))
	packed.Lsh(packed, 64)
	packed.Or(packed, big.NewInt(int64(a)))

	return packed
}
