package main

import (
	"fmt"
	"net/http"
	"strconv"
	"log"
)

func main() {
	http.HandleFunc("/", inc)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func inc(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("counter")
	if err != nil {
		c = &http.Cookie{
			Name:  "counter",
			Value: "1",
			Path:  "/",
		}
	} else {
		count, err := strconv.Atoi(c.Value)
		if err != nil {
			log.Fatalln(err)
		}
		count++
		c.Value = strconv.Itoa(count)
	}

	http.SetCookie(w, c)
	fmt.Fprintln(w, "YOUR COUNTER:", c.Value)
}
