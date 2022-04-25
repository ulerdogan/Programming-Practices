package main

import (
	"io"
	"net/http"
)

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "ulas")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func _init(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world")
}

func main() {
	http.HandleFunc("/", _init)
	http.HandleFunc("/me/", me)
	http.HandleFunc("/dog/", dog)

	http.ListenAndServe(":8080", nil)
}