package handlers

import(
	"books-api/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)


// In-memory storage for our books
// In a real application, this would be a database
var books []models.Book
var nextID = 1 // Simple ID counter


// GetBooks handles GET /books
// Returns all books as JSON
func GetBooks(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to tell client we're sending JSON
	w.Header().Set("Content-Type", "application/json")
	
	// json.NewEncoder creates an encoder that writes to w (the HTTP response)
	// Encode converts our books slice to JSON and writes it to the response
	json.NewEncoder(w).Encode(books)
}

// GetBook handles GET /books/{id}
// Returns a single book by ID
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Extract the ID from the URL path
	// URL like "/books/123" becomes ["", "books", "123"] when split by "/"
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	
	// Convert the string ID to integer
	// pathParts[2] is "123" in our example
	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	
	// Search through our books slice to find matching ID
	for _, book := range books {
		if book.ID == id {
			// Found it! Send back the book as JSON
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	
	// If we reach here, no book was found with that ID
	http.Error(w, "Book not found", http.StatusNotFound)
}

// CreateBook handles POST /books
// Creates a new book from JSON in request body
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Only allow POST method for this handler
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Create an empty book struct to decode JSON into
	var book models.Book
	
	// json.NewDecoder reads from r.Body (the request body)
	// Decode converts JSON to our Go struct
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	// Set the ID and creation timestamp
	book.ID = nextID
	nextID++ // Increment for next book
	book.CreatedAt = time.Now()
	
	// Add the new book to our in-memory storage
	books = append(books, book)
	
	// Return the created book with 201 Created status
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// UpdateBook handles PUT /books/{id}
// Updates an existing book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Extract ID from URL (same as GetBook)
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	
	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	
	// Parse the updated book data from request body
	var updatedBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	// Find the book to update
	for i, book := range books {
		if book.ID == id {
			// Preserve the original ID and CreatedAt timestamp
			updatedBook.ID = book.ID
			updatedBook.CreatedAt = book.CreatedAt
			
			// Replace the book in our slice
			books[i] = updatedBook
			
			// Return the updated book
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	
	// Book with that ID doesn't exist
	http.Error(w, "Book not found", http.StatusNotFound)
}

// DeleteBook handles DELETE /books/{id}
// Removes a book from storage
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Extract ID from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	
	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	
	// Find and remove the book
	for i, book := range books {
		if book.ID == id {
			// Remove book from slice by combining parts before and after index i
			// books[:i] = everything before index i
			// books[i+1:] = everything after index i
			books = append(books[:i], books[i+1:]...)
			
			// Return 204 No Content status (successful deletion, no body)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	
	// Book not found
	http.Error(w, "Book not found", http.StatusNotFound)
}