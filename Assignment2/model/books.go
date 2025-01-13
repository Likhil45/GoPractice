package model

import "fmt"

type Books struct {
	Title           string
	Author          string
	Pages           int
	CopiesAvailable int
}

func (b Books) Display() {
	fmt.Printf("Book: \n\tBook Name: %s\n\tAuthor: %s\n\tPages: %d\n\tCopies Available: %d\n", b.Title, b.Author, b.Pages, b.CopiesAvailable)

}
func CreateBook(Title string, Author string, Pages int, Copies int) *Books {
	book := Books{Title: Title, Author: Author, Pages: Pages, CopiesAvailable: Copies}
	return &book
}

func (b *Books) Borrow() {
	if b.CopiesAvailable > 0 {
		b.CopiesAvailable--
		fmt.Println("Successfuly borrowed the book")
	} else {
		fmt.Println("No books are left")
	}
}

func (b *Books) ReturnBook() {
	fmt.Println("Book returned successfuly")
	b.CopiesAvailable++
}

func SwapTitles(a *Books, b *Books) {
	var temp string = a.Title
	a.Title = b.Title
	b.Title = temp
	fmt.Println("Successfuly Swapped the titles")
	a.Display()
	b.Display()
}
