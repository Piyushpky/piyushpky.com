package main

import (
	"fmt"
	"net/http"
	"os" // Import the os package
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from Vercel Go Server!")
}

func main() {
	http.HandleFunc("/", handler)

	// Get port from environment, default to 8080
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080" // Default port if not set
	}
	port, _ := strconv.Atoi(portStr)

	fmt.Printf("Server listening on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
