func Home(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "home.page.tmpl")
}

func About(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "about.page.tmpl")
}

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