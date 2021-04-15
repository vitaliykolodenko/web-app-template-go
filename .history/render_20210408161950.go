package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(rw http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		fmt.Println("failed to parse template: " + tmpl)
	}
	err = t.Execute(rw, nil)

	if err != nil {
		fmt.Println("Cannot render template")
	}
}
