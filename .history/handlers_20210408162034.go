package main

import (
	"net/http"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "home.page.tmpl")
}

func About(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "about.page.tmpl")
}
