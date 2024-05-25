package book

import "fmt"

// if you want to convert an attribute or method like private, you can use lowercase

// Book is a struct that represents a book
type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

type Book struct {
	title  string
	author string
	pages  int
}

// methods

// NewBook creates a new Book
func NewBook(title, author string, pages int) *Book {
	return &Book{title, author, pages}
}

// PrintInfo prints the book information
func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", b.title, b.author, b.pages)
}

// setters

// SetTitle sets the title of the book
func (b *Book) SetTitle(title string) {
	b.title = title
}

// SetAuthor sets the author of the book
func (b *Book) SetAuthor(author string) {
	b.author = author
}

// SetPages sets the number of pages of the book
func (b *Book) SetPages(pages int) {
	b.pages = pages
}

// getters

// GetTitle returns the title of the book
func (b *Book) GetTitle() string {
	return b.title
}

// GetAuthor returns the author of the book
func (b *Book) GetAuthor() string {
	return b.author
}

// GetPages returns the number of pages of the book
func (b *Book) GetPages() int {
	return b.pages
}

// Composition

type TextBook struct {
	Book
	edition int
	level   string
}

// NewTextBook creates a new TextBook
func NewTextBook(title, author string, pages, edition int, level string) *TextBook {
	return &TextBook{
		Book:    Book{title, author, pages},
		edition: edition,
		level:   level,
	}
}

// PrintInfo prints the book information
func (b *TextBook) PrintInfo() {
	b.Book.PrintInfo()
	fmt.Printf("Edition: %d\nLevel: %s\n", b.edition, b.level)
}
