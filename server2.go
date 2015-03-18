package main

import (
	"net/http"
	"text/template"
)

type Page struct {
	Title string
	Num   int
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"Hello World.", 1}
	tmpl, err := template.New("new").Parse("Title = {{.Title}}, num = {{.Num}}")
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
