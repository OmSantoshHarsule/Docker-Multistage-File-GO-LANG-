package main

import (
	"fmt"
	"net/http"
)

// handler function that writes "Hello, World!" to the response
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Register the handler function for the "/" route
	http.HandleFunc("/", helloHandler)

	// Define the port to listen on
	port := ":3000"
	fmt.Println("Starting server on port", port)

	// Start the server and listen on the defined port
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
