package main

type Book struct {
}

func (*Book) IsBookTaken() bool { return false }

func (*Book) ReadBook() []string { return nil }

type BookSearcher struct {
	byTitle  map[string]*Book
	byAuthor map[string]*Book
	byId     map[string]*Book
}
