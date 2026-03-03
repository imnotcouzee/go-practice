package warehouse

import "fmt"

type supplier struct {
	id   int
	name string
}

func newSupplier(id int, name string) *supplier {
	if id <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : ID должен быть больше нуля")
		return nil
	}

	if name == "" {
		fmt.Println("")
		fmt.Println("Ошибка : Имя должно быть заполнено")
		return nil
	}

	return &supplier{id: id, name: name}
}
