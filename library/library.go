package library

import (
	"fmt"
	"library/book"
)

type Library struct {
	Name  string
	Books map[string]*book.Book
}

func NewLibrary(name string) *Library {
	if name == "" {
		fmt.Println("")
		fmt.Println("Ошибка : Имя должно быть заполнено")
		return nil
	}

	return &Library{Name: name, Books: make(map[string]*book.Book)}
}

func (l *Library) AddBook(title, author string, copies int) {
	if _, ok := l.Books[title]; ok {
		l.Books[title].Copies += copies
		return
	}

	l.Books[title] = book.NewBook(title, author, copies)
}

func (l *Library) GetBook(title string) *book.Book {
	if title == "" {
		fmt.Println("")
		fmt.Println("Ошибка : Название книги должно быть заполнено")
		return nil
	}

	if _, ok := l.Books[title]; !ok {
		fmt.Println("")
		fmt.Println("Ошибка : Искомая книга не найдена")
		return nil
	}

	return l.Books[title]
}

func (l *Library) RemoveBook(title string, copies int) {
	if title == "" {
		fmt.Println("")
		fmt.Println("Ошибка : Название книги должно быть заполнено")
		return
	}

	if copies <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : Количество копий для удаления должно быть больше нуля")
		return
	}

	if _, ok := l.Books[title]; !ok {
		fmt.Println("")
		fmt.Println("Ошибка : Книга для удаления не найдена")
		return
	}

	l.GetBook(title).Copies -= copies

	if l.GetBook(title).Copies <= 0 {
		l.GetBook(title).Copies = 0
		delete(l.Books, title)
	}

}

func (l *Library) ListBooks() {
	if len(l.Books) <= 0 {
		fmt.Println("")
		fmt.Println("Нет ни одной книги для вывода")
		return
	}

	fmt.Println("")
	fmt.Println("Список всех книг в библиотеке:")

	for _, b := range l.Books {
		fmt.Println("Название:", b.Title, "; Автор:", b.Author, "; Количество копий:", b.Copies)
	}

}
