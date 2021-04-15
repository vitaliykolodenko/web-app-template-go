package main

import (
	"fmt"
	"net/http"

	"github.com/vitaliykolodenko/go-course/pkg/handlers"
)

func main() {
	fmt.Println("Hello, go!")
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	http.ListenAndServe(":9090", nil)
}
