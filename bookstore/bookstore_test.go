package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Ensure we have a Book struct
func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		ID:     0,
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

// Test the buying behavior which takes some
// Book value, decrements its number of copies
// by 1 and returns the modified Book value
func TestBuy(t *testing.T) {
	t.Parallel()
	testBook := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
	want := 1
	result, err := bookstore.Buy(testBook)
	got := result.Copies
	if err != nil {
		t.Fatalf("want no error for valid input, got %v", err)
	}
	if got != want {
		t.Errorf("Expecting %v copies, got %v copies instead", want, got)
	}
}

// Test the buying behavior which takes some
// Book with no copies and returns an error
func TestBuyZeroCopies(t *testing.T) {
	t.Parallel()
	testBook := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}
	_, err := bookstore.Buy(testBook)
	if err == nil {
		t.Error("Expected error, but got err == nil")
	}
}

// Test the behavior to obtain a list of books
func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}
	got := bookstore.GetAllBooks(catalog)
	// Introducing the cmp package here
	// Use Equal() for comparing structs or slices
	// Use Diff() for printing the line by line diff
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	testCatalog := []bookstore.Book{
		{ID: 0, Title: "For the Love of Go"},
		{ID: 1, Title: "Azkaban the Prisoner Guy"},
	}
	got, err := bookstore.GetBook(testCatalog, 1)
	want := testCatalog[1]
	if err != nil {
		t.Fatalf("expected nil error for valid input, got %v", err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookNoID(t *testing.T) {
	t.Parallel()
	testCatalog := []bookstore.Book{
		{ID: 0, Title: "For the Love of Go"},
		{ID: 1, Title: "Azkaban the Prisoner Guy"},
	}
	_, err := bookstore.GetBook(testCatalog, 2)
	if err == nil {
		t.Errorf("Expected error, got err == nil")
	}
}
