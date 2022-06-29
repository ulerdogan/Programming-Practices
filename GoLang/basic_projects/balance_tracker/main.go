package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"balance_tracker/util"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

var tpl *template.Template
var client *ethclient.Client

type info struct {
	Address string
	// TODO: Add query times and make lower q's
}

type blnc struct {
	Amount decimal.Decimal
	Price  decimal.Decimal
}

func init() {
	var err error
	client, err = ethclient.Dial("https://api.avax.network/ext/bc/C/rpc")

	if err != nil {
		log.Fatalln("Network connection error.")
	}

	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/balance", balance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

	fmt.Println("we have a connection", client)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("account")
	var i info

	if err != http.ErrNoCookie {
		i = info{cookie.Value}
	}

	if req.Method == http.MethodPost {
		addr := req.FormValue("address")

		if util.IsValidAddress(addr) {
			cookie = &http.Cookie{
				Name:  "account",
				Value: addr,
				Path:  "/",
			}

			i = info{addr}
		}
	}

	http.SetCookie(w, cookie)
	tpl.ExecuteTemplate(w, "index.gohtml", i)
}

func balance(w http.ResponseWriter, req *http.Request) {
	if !isSubmitted(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	cookie, _ := req.Cookie("account")

	account := common.HexToAddress(cookie.Value)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: 15 is hardcoded - change it by a price api
	b := blnc{util.ToDecimal(balance, 18), util.ToDecimal(balance, 18).Mul(decimal.NewFromInt(15))}

	tpl.ExecuteTemplate(w, "balance.gohtml", b)
}

func isSubmitted(req *http.Request) bool {
	cookie, err := req.Cookie("account")

	if err == http.ErrNoCookie {
		return false
	}

	if util.IsValidAddress(cookie.Value) {
		return true
	}

	return false
}
