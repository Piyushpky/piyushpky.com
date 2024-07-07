package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func hHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! first page")
}

func main() {
	http.HandleFunc("/api", helloHandler)
	http.HandleFunc("/", hHandler)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
