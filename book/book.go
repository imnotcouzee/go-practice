package book

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

func NewBook(title, author string, copies int) *Book {
	if title == "" {
		fmt.Println("")
		fmt.Println("Ошибка : Название книги должно быть заполнено")
		return nil
	}

	if author == "" {
		fmt.Println("")
		fmt.Println("Ошибка : Имя автора должно быть заполнено")
		return nil
	}

	if copies < 0 {
		fmt.Println("")
		fmt.Println("Ошибка : Количество копий не должно быть отрицательное")
		return nil
	}

	return &Book{Title: title, Author: author, Copies: copies}
}
