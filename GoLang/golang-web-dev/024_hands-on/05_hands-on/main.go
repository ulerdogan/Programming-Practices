package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("starting-files/templates/index.gohtml"))
}

func main() {
	http.Handle("/pics/", http.FileServer(http.Dir("starting-files/public")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.Execute(w, nil)
		if err != nil {
			log.Fatalln("err:", err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}