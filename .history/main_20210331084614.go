package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, go!")
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(rw, "Hello world")
		if err != nil {
			fmt.Printf("Erorr happened: %s", err)
		} else {
			fmt.Printf("Send number of byes: %d", n)
		}
	})

	http.ListenAndServe(":9090", nil)
}
