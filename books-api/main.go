
package main

import (
	"books-api/handlers"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// router function handles all incoming HTTP requests
// It looks at the URL path and HTTP method to decide which handler to call
func router(w http.ResponseWriter, r *http.Request) {
	// Print request info for debugging
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)
	
	// Route requests based on URL path and HTTP method
	switch {
	// GET /books - get all books
	case r.URL.Path == "/books" && r.Method == http.MethodGet:
		handlers.GetBooks(w, r)
	
	// POST /books - create new book
	case r.URL.Path == "/books" && r.Method == http.MethodPost:
		handlers.CreateBook(w, r)
	
	// GET /books/{id} - get specific book
	case strings.HasPrefix(r.URL.Path, "/books/") && r.Method == http.MethodGet:
		handlers.GetBook(w, r)
	
	// PUT /books/{id} - update specific book
	case strings.HasPrefix(r.URL.Path, "/books/") && r.Method == http.MethodPut:
		handlers.UpdateBook(w, r)
	
	// DELETE /books/{id} - delete specific book
	case strings.HasPrefix(r.URL.Path, "/books/") && r.Method == http.MethodDelete:
		handlers.DeleteBook(w, r)
	
	// No matching route found
	default:
		http.Error(w, "Not Found check your URL", http.StatusNotFound)
	}
}

func main() {
	// Register our router function to handle all HTTP requests
	// The "/" pattern matches all paths
	http.HandleFunc("/", router)
	
	// Print startup information
	fmt.Println("📚 Books API Server")
	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("\nAvailable endpoints:")
	fmt.Println("  GET    /books      - Get all books")
	fmt.Println("  POST   /books      - Create a new book")
	fmt.Println("  GET    /books/{id} - Get a specific book")
	fmt.Println("  PUT    /books/{id} - Update a specific book")
	fmt.Println("  DELETE /books/{id} - Delete a specific book")
	fmt.Println("\nPress Ctrl+C to stop the server")
	
	// Start the HTTP server on port 8080
	// log.Fatal will print any error and exit the program if server fails to start
	log.Fatal(http.ListenAndServe(":8080", nil))
}