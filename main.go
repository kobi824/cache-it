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

func main() {
	fmt.Println("If a book exists in the database use the cache to eliminate db calls")
	fmt.Println("using test to validate program")
}
