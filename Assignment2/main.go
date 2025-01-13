package main

import (
	"practice.com/GoPractice/Assignment2/model"
)

func main() {
	book := model.CreateBook("Wings of Fire", "APJ Abdul Kalaam", 120, 100)
	book1 := model.CreateBook("48 laws of Power", "Town Hill", 200, 15)
	book.Display()
	book1.Display()
	book.Borrow()
	book.ReturnBook()
	model.SwapTitles(book, book1)
}
