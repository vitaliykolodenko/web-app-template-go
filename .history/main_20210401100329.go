package main

import (
	"fmt"
	"net/http"
)

func Home(rw http.ResponseWriter, r *http.Request) {

}

func About(rw http.ResponseWriter, r *http.Request) {

}

func main() {
	fmt.Println("Hello, go!")
	http.HandleFunc("/", Home)

	http.ListenAndServe(":9090", nil)
}
