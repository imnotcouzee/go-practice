package main

import (
	contactbook "contact-book/contact-book"
)

func main() {
	myContactBook := contactbook.NewContactBook("couzee")

	myContactBook.AddContact("vasya", "123")
	myContactBook.AddContact("couzee", "1234")

	myContactBook.List()

	myContactBook.Delete("vasya")

	myContactBook.List()
}
