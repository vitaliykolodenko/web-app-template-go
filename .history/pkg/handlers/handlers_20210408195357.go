package handlers

import (
	"net/http"

	"github.com/vitaliykolodenko/go-course/pkg/render"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "home.page.tmpl")
}

func About(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "about.page.tmpl")
}
