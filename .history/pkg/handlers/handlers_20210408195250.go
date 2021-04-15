package handlers

import (
	"net/http"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "home.page.tmpl")
}

func About(rw http.ResponseWriter, r *http.Request) {
	RenderTemplate(rw, "about.page.tmpl")
}
