package bookstore

import "errors"

// Storing information about books
type Book struct {
	Title  string
	Author string
	Copies int
}

// Buying a book decrements its number of copies by 1
func Buy(b Book) (Book, error) {
	if b.Copies > 0 {
		b.Copies -= 1
		return b, nil
	} else {
		return Book{}, errors.New("cannot buy a book with no copies")
	}
}
