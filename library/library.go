package library

import "fmt"

type Library struct {
	Name    string
	Books   map[int]*Book
	Readers map[int]*Reader
	Borrows map[int]*Borrow
}

func NewLibrary(name string) *Library {
	if name == "" {
		fmt.Println("")
		fmt.Println("Ошибка : Название библиотеки не должно быть пустое")
		return nil
	}

	return &Library{Name: name, Books: make(map[int]*Book), Readers: make(map[int]*Reader), Borrows: make(map[int]*Borrow)}
}

func (l *Library) AddBook(book *Book) {
	if book == nil {
		fmt.Println("")
		fmt.Println("Ошибка : Книга не найдена")
		return
	}

	if _, ok := l.Books[book.ID]; ok {
		fmt.Println("")
		fmt.Println("Ошибка : Книга с таким ID уже зарегистрирована в библиотеке")
		return
	}

	l.Books[book.ID] = book
}

func (l *Library) RegisterReader(reader *Reader) {
	if reader == nil {
		fmt.Println("")
		fmt.Println("Ошибка : Читатель не найден")
		return
	}

	if _, ok := l.Readers[reader.ID]; ok {
		fmt.Println("")
		fmt.Println("Ошибка : Читатель с таким ID уже зарегестрирован в системе")
		return
	}

	l.Readers[reader.ID] = reader
}

func (l *Library) BorrowBook(readerID, bookID, borrowID int) {
	if readerID <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : ID читателя должен быть больше нуля")
		return
	}

	if bookID <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : ID книги должен быть больше нуля")
		return
	}

	if borrowID <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : ID выдачи должен быть больше нуля")
		return
	}

	if _, ok := l.Readers[readerID]; !ok {
		fmt.Println("")
		fmt.Println("Ошибка : Читатель с таким ID не найден в системе")
		return
	}

	findReader := l.Readers[readerID]

	if findReader == nil {
		fmt.Println("")
		fmt.Println("Ошибка : При выборе читателя произошла ошибка")
		return
	}

	if len(findReader.ListBorrows) >= 3 {
		fmt.Println("")
		fmt.Println("Ошибка : Нельзя взять больше 3 книг")
		return
	}

	if _, ok := l.Books[bookID]; !ok {
		fmt.Println("")
		fmt.Println("Ошибка : Книга с таким ID не найден в системе")
		return
	}

	findBook := l.Books[bookID]

	if findBook == nil {
		fmt.Println("")
		fmt.Println("Ошибка : При выборе книги произошла ошибка")
		return
	}

	if _, ok := l.Borrows[borrowID]; ok {
		fmt.Println("")
		fmt.Println("Ошибка : Выдача с таким ID уже есть в системе")
		return
	}

	if findBook.Available <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : Нет доступной книги")
		return
	}

	newBorr := NewBorrow(borrowID, findReader, findBook)

	if newBorr == nil {
		fmt.Println("")
		fmt.Println("Ошибка : При создании книги произошла ошибка")
		return
	}

	l.Borrows[borrowID] = newBorr

	findReader.ListBorrows = append(findReader.ListBorrows, newBorr)

	findBook.Available--

	if findBook.Available < 0 {
		findBook.Available = 0
	}

}

func (l *Library) ReturnBook(borrowID int) {
	if borrowID <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : ID выдачи должен быть больше нуля")
		return
	}

	if _, ok := l.Borrows[borrowID]; !ok {
		fmt.Println("")
		fmt.Println("Ошибка : Выдача с таким ID не найдена")
		return
	}

	if l.Borrows[borrowID].Status != "Активна" {
		fmt.Println("")
		fmt.Println("Ошибка: Книга и так не занята")
		return
	}

	l.Borrows[borrowID].Book.Available++
	l.Borrows[borrowID].Status = "Возвращена"

}

func (l *Library) GetAvailableBooks() {
	var haveBook bool

	fmt.Println("")
	fmt.Println("Список доступных книг:")

	for _, book := range l.Books {
		if book.Available > 0 {
			fmt.Println("ID -", book.ID, ";", book.Title, "; Доступно -", book.Available)
			haveBook = true
		}
	}

	if !haveBook {
		fmt.Println("Доступных книг нет")
		return
	}

}

func (l *Library) GetReaderBorrows(readedID int) {
	fmt.Println("")
	fmt.Println("Список выдачей читателя:")

	if readedID <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : ID должен быть больше нуля")
		return
	}

	for _, reader := range l.Readers {
		if readedID == reader.ID {
			for _, borrows := range reader.ListBorrows {
				fmt.Println("-----------")
				fmt.Println("ID выдачи:", borrows.ID)
				fmt.Println("ID книги:", borrows.Book.ID, "|", "Автор:", borrows.Book.Author, "|", "Название:", borrows.Book.Title)
			}
		}
	}

}
