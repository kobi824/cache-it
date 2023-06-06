package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	for i, _ := range [1000]int{} {
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

func (s *Server) GetBook(w http.ResponseWriter, r *http.Request) {
	urlId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(urlId)

	book, ok := s.useCache(id)
	if ok {
		json.NewEncoder(w).Encode(book)
		return
	}

	book, ok = s.db[id]
	if !ok {
		log.Printf("given id does not return a book %v", s.db[id])
	}
	s.count++

	s.cache[id] = book
	json.NewEncoder(w).Encode(book)
}

func (s *Server) useCache(id int) (*Book, bool) {
	book, ok := s.cache[id]
	return book, ok
}

func main() {
	fmt.Println("If a book exists in the database use the cache to eliminate db calls")
	fmt.Println("using test to validate program")
}
