package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"balance_tracker/util"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

var tpl *template.Template
var client *ethclient.Client

type info struct {
	Address string
}

type blnc struct {
	Amount decimal.Decimal
	Price  decimal.Decimal
	Total  decimal.Decimal
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

	cookieT, err := req.Cookie("timestamp")

	if err == http.ErrNoCookie {
		cookieT = &http.Cookie{
			Name:  "timestamp",
			Value: "0",
			Path:  "/",
		}
	}

	var b blnc

	cookieB, err := req.Cookie("balance")
	if err == http.ErrNoCookie {
		cookieB = &http.Cookie{
			Name:  "balance",
			Value: time.UnixDate,
			Path:  "/",
		}
		cookieT.Value = time.UnixDate
	}

	ct, err := time.Parse(time.UnixDate, cookieT.Value)
	if err != nil {
		log.Println(err)
	}

	if time.Since(ct) > (time.Second * 120) {
		cookieAcc, _ := req.Cookie("account")

		balance := getBalance(common.HexToAddress(cookieAcc.Value))

		b = blnc{balance, getPrice(), decimal.Zero}
		b.Total = balance.Mul(getPrice())
		cookieT.Value = time.Now().Format(time.UnixDate)

		cookieB.Value = balance.String()
	} else {
		cb, _ := decimal.NewFromString(cookieB.Value)
		b = blnc{cb, getPrice(), decimal.Zero}
		b.Total = b.Amount.Mul(getPrice())
	}

	http.SetCookie(w, cookieT)
	http.SetCookie(w, cookieB)
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