package bookstore

import "errors"

// Storing information about books
type Book struct {
	ID     int
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

// Get the catalog of books
func GetAllBooks(catalog []Book) []Book {
	return catalog
}

// Get a specific book
func GetBook(catalog []Book, ID int) (Book, error) {
	for _, b := range catalog {
		if b.ID == ID {
			return b, nil
		}
	}
	return Book{}, errors.New("book not found")
}
