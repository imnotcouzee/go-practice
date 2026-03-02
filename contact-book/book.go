package contactbook

import (
	"fmt"
)

type ContactBook struct {
	name     string
	contacts map[string]string
}

func NewContactBook(name string) *ContactBook {
	if name == "" {
		fmt.Println("")
		fmt.Println("Название книги контактов должно быть заполнено")
		return nil
	}

	return &ContactBook{name: name, contacts: make(map[string]string)}

}
func (c *ContactBook) AddContact(name string, phone string) {
	if name == "" {
		fmt.Println("")
		fmt.Println("Ошибка : Имя должно быть заполнено")
		return
	}

	if phone == "" {
		fmt.Println("")
		fmt.Println("Ошибка : Телефон должен быть заполнен")
		return
	}

	c.contacts[name] = phone
}

func (c *ContactBook) GetPhone(name string) string {
	if name == "" {
		fmt.Println("Ошибка : Имя должно быть заполнено")
		return ""
	}

	if value, ok := c.contacts[name]; !ok {
		fmt.Println("Ошибка : Для данного имени номера телефона не найдено")
		return ""
	} else {
		fmt.Println("Номер телефона:", value)
		return value
	}
}

func (c *ContactBook) Delete(name string) {
	if name == "" {
		fmt.Println("Ошибка : Имя должно быть заполнено")
		return
	}

	delete(c.contacts, name)
}

func (c *ContactBook) List() {
	fmt.Println("")
	fmt.Println("Список телефонной книги:")
	for name, phone := range c.contacts {
		fmt.Println("Имя:", name, "Номер телефона:", phone)
	}
}
