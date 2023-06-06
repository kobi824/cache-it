package main

import (
	"fmt"
)

type Book struct {
	ID    int
	Title string
}

type Server struct {
	db    map[int]*Book
	count int
	cache map[int]*Book
}

func NewServer() *Server {
	db := make(map[int]*Book)
	for i, _ := range [4000]int{} {
		db[i+1] = &Book{
			ID:    i + 1,
			Title: fmt.Sprintf("Book-%d", i+1),
		}
	}
	return &Server{
		db:    db,
		cache: make(map[int]*Book),
	}
}

func main() {
	fmt.Println("If a book exists in the database use the cache to eliminate db calls")
	fmt.Println("using test to validate program")
}
