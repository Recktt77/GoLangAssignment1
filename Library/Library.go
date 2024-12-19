package Library

import (
	"fmt"
)

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]*Book
}

func (l *Library) init() {
	if l.Books == nil {
		l.Books = make(map[string]*Book)
	}
}

func (l *Library) AddBook(ID string, Title string, Author string) {
	l.Books[ID] = &Book{
		ID:         ID,
		Title:      Title,
		Author:     Author,
		IsBorrowed: false,
	}

	fmt.Printf("book '%s' added\n", Title)
}

func (l *Library) BorrowBook(id string) {
	_, ok := l.Books[id]
	if !ok {
		fmt.Printf("Book doesn't exist")
	}
	l.Books[id].IsBorrowed = true
	fmt.Printf("Book #%v borrowed\n", id)
}
func (l *Library) ReturnBook(id string) {
	l.Books[id].IsBorrowed = false

	fmt.Printf("Book returned")
}

func (l *Library) ListBooks() {
	for id, book := range l.Books {
		fmt.Printf("Id: %v book: %v\n", id, book)
	}
}
