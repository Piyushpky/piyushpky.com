package main

import (
	"fmt"
	"github.com/piyushpky/piyushpky.com/api"
	"net/http"
	"os" // Import the os package
	"strconv"
)

func main() {
	http.HandleFunc("/", api.Handler)

	// Get port from environment, default to 8080
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080" // Default port if not set
	}
	port, _ := strconv.Atoi(portStr)

	fmt.Printf("Server listening on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
