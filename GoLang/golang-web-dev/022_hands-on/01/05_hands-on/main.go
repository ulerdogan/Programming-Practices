package main

import (
	"html/template"
	"log"
	"io"
	"net/http"
)

func me(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("exp.gohtml")
	if err != nil {
		log.Fatalln("error", err)
	}

	err = tpl.ExecuteTemplate(w, "exp.gohtml", "ulas")
	if err != nil {
		log.Fatalln("error", err)
	}
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func _init(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world")
}

func main() {
	http.Handle("/", http.HandlerFunc(_init))
	http.Handle("/me/", http.HandlerFunc(me))
	http.Handle("/dog/", http.HandlerFunc(dog))

	http.ListenAndServe(":8080", nil)
}