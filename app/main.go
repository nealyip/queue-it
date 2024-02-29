package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a new instance of the HTTP router (ServeMux)
	mux := http.NewServeMux()

	// Define the root handler function
	rootHandler := func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`<a href=/cart>Add an item to Cart</a>`))
		if err != nil {
			return
		}
	}

	// Define the cart handler function
	cartHandler := func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`<p>Item added to cart!</p>`))
		if err != nil {
			return
		}
	}

	// Register the middleware and handlers to their respective routes
	mux.HandleFunc("/", Middleware(rootHandler))
	mux.HandleFunc("/cart", Middleware(cartHandler))

	// Start the HTTP server
	log.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
