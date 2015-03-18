package main

import (
	"net/http"
	"strconv"
	"text/template"
)

type Page struct {
	Title string
	Num   int
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	num, _ := strconv.Atoi(r.FormValue("num"))
	page := Page{"Hello World.", num}
	tmpl, err := template.ParseFiles("sample.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
