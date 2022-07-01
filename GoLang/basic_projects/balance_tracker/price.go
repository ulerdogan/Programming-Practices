package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"context"
	"balance_tracker/util"
	"github.com/shopspring/decimal"
	"github.com/ethereum/go-ethereum/common"
)

type Token struct {
	Price struct {
		Usd decimal.Decimal `json:"usd"`
	} `json:"avalanche-2"`
}

var lastPrice decimal.Decimal
var lastCheck time.Time

func getPrice() decimal.Decimal {

	if time.Now().Sub(lastCheck) > (time.Second * 60) {

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

		lastPrice = responseObject.Price.Usd
		lastCheck = time.Now()
	}

	return lastPrice
}

func getBalance(account common.Address) decimal.Decimal {

	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	return util.ToDecimal(balance, 18)
}
