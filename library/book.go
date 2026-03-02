package library

import "fmt"

type Book struct {
	ID        int
	Title     string
	Author    string
	Copies    int
	Available int
}

func NewBook(id int, title, author string, copies, available int) *Book {
	if id <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : ID книги должен быть больше нуля")
		return nil
	}

	if title == "" {
		fmt.Println("")
		fmt.Println("Ошибка : Название книги не должно быть пустое")
		return nil
	}

	if author == "" {
		fmt.Println("")
		fmt.Println("Ошибка : Имя автора не должно быть пустое")
		return nil
	}

	if copies < 0 {
		fmt.Println("")
		fmt.Println("Ошибка : Количество экземпляров книг не должно быть отрицательное")
		return nil
	}

	if available < 0 {
		fmt.Println("")
		fmt.Println("Ошибка : Количество доступных книг не должно быть отрицательное")
		return nil
	}

	if available > copies {
		fmt.Println("")
		fmt.Println("Ошибка : Количество доступных книг не должно превышать общее количество экземпляров")
		return nil
	}

	return &Book{ID: id, Title: title, Author: author, Copies: copies, Available: available}
}
