package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, go!")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(":9090", nil)
}
