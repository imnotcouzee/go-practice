package library

import "fmt"

type Borrow struct {
	ID     int
	Reader *Reader
	Book   *Book
	Status string
}

func NewBorrow(id int, reader *Reader, book *Book) *Borrow {
	if id <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : ID выдачи должен быть больше нуля")
		return nil
	}

	if reader == nil {
		fmt.Println("")
		fmt.Println("Ошибка : Читатель не найден")
		return nil
	}

	if book == nil {
		fmt.Println("")
		fmt.Println("Книга не найдена")
		return nil
	}

	return &Borrow{ID: id, Reader: reader, Book: book, Status: "Активна"}
}
