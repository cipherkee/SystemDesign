package main

type OnlineBookReader struct {
	bookSearcher *BookSearcher

	bookReaderMap map[*Book]*Reader
	readerBookMap map[*Reader]*Book
}

func (o *OnlineBookReader) SearchBook() []*Book { return nil }

func (o *OnlineBookReader) AssignABookForReader(*Reader, *Book) {
}
