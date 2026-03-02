package main

import (
	"library-system/library"
)

func main() {
	centralLibrary := library.NewLibrary("Центральная Библиотека")

	firstBook := library.NewBook(1, "Война и Мир. Том 1", "Толстой Л.Н.", 10, 5)
	secondBook := library.NewBook(2, "Война и Мир. Том 2", "Толстой Л.Н.", 10, 5)

	firstReader := library.NewReader(1, "vasya")
	secondReader := library.NewReader(2, "oleg")

	centralLibrary.AddBook(firstBook)
	centralLibrary.AddBook(secondBook)

	centralLibrary.RegisterReader(firstReader)
	centralLibrary.RegisterReader(secondReader)

	centralLibrary.BorrowBook(1, 1, 1)

	centralLibrary.GetReaderBorrows(1)

	centralLibrary.ReturnBook(1)

	centralLibrary.GetReaderBorrows(1)

	centralLibrary.BorrowBook(1, 1, 2)

	centralLibrary.GetReaderBorrows(1)

	centralLibrary.ReturnBook(3)

}
