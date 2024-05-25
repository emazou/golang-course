package main

import (
	"fmt"
	"library/animal"
	"library/book"
)

func main() {
	myBook := book.NewBook(
		"The Art of Computer Programming",
		"Donald Knuth",
		700,
	)
	myBook.PrintInfo()
	myBook.SetTitle("The Go Programming Language")
	myBook.SetPages(200)
	myBook.PrintInfo()
	fmt.Println(myBook.GetTitle())

	myTextBook := book.NewTextBook(
		"Introduction to Algorithms",
		"Thomas H. Cormen",
		800,
		3,
		"Undergraduate",
	)
	myTextBook.PrintInfo()

	book.Print(myBook)
	book.Print(myTextBook)

	// animals

	myDog := animal.Dog{Name: "Buddy"}
	myCat := animal.Cat{Name: "Lucy"}

	animals := []animal.Animal{
		&myDog,
		&myCat,
		&animal.Dog{Name: "Max"},
		&animal.Cat{Name: "Luna"},
	}

	for _, a := range animals {
		animal.PrintSoundAnimal(a)
	}
}
