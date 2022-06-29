package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/shopspring/decimal"
)

type Token struct {
	Price struct {
		Usd decimal.Decimal `json:"usd"`
	} `json:"avalanche-2"`
}

func getPrice() decimal.Decimal {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=avalanche-2&vs_currencies=usd")
	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var responseObject Token
	json.Unmarshal(body, &responseObject)
	return responseObject.Price.Usd
}
