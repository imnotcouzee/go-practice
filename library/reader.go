package library

import (
	"fmt"
)

type Reader struct {
	ID          int
	Name        string
	ListBorrows []*Borrow
}

func NewReader(id int, name string) *Reader {
	if id <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : ID читателя должен быть больше нуля")
		return nil
	}

	if name == "" {
		fmt.Println("")
		fmt.Println("Ошибка : Имя читателя не должно быть пустое")
		return nil
	}

	return &Reader{ID: id, Name: name}
}
