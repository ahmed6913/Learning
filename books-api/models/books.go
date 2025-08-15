package models

import"time"

type Book struct{

	ID          int       `json:"id"`           // Will appear as "id" in JSON
	Title       string    `json:"title"`        // Will appear as "title" in JSON
	Author      string    `json:"author"`       // Will appear as "author" in JSON
	ISBN        string    `json:"isbn"`         // Will appear as "isbn" in JSON
	PublishedAt time.Time `json:"published_at"` // Will appear as "published_at" in JSON
	CreatedAt   time.Time `json:"created_at"`   // Will appear as "created_at" in JSON

}

// NewBook is a constructor function that creates a new Book
// It takes the required fields and sets CreatedAt to current time
// Returns a pointer to Book (*Book) for efficiency
func NewBook(title, author, isbn string, publishedAt time.Time) *Book {
	return &Book{
		Title:       title,
		Author:      author,
		ISBN:        isbn,
		PublishedAt: publishedAt,
		CreatedAt:   time.Now(), // Set creation time to now
	}
}