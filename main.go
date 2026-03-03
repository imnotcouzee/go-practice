package main

import (
	"library/library"
)

func main() {

	centralLibrary := library.NewLibrary("Центральная Библиотека")

	centralLibrary.AddBook("Гарри Поттер", "Дж.К. Роулинг", 3)
	centralLibrary.AddBook("Властелин колец", "Дж.Р.Р. Толкин", 2)
	centralLibrary.AddBook("Гарри Поттер", "Дж.К. Роулинг", 1)

	centralLibrary.ListBooks()

	centralLibrary.RemoveBook("Гарри Поттер", 2)

	centralLibrary.ListBooks()

	centralLibrary.RemoveBook("Гарри Поттер", 2)

	centralLibrary.ListBooks()
}
