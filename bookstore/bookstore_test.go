package bookstore_test

import (
	"bookstore"
	"testing"
)

// Ensure we have a Book struct
func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
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
