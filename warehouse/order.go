package warehouse

import "fmt"

const (
	CREATED  = "Создан"
	ACCEPTED = "Подтвержден"
	CANCELED = "Отменен"
)

type order struct {
	id       int
	supplier *supplier
	products map[*product]int
	status   string
}

func newOrder(id int, supplier *supplier) *order {
	if id <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : ID должен быть больше нуля")
		return nil
	}

	if supplier == nil {
		fmt.Println("")
		fmt.Println("Ошибка : Поставщик не найден")
		return nil
	}

	return &order{id: id, supplier: supplier, products: make(map[*product]int), status: CREATED}

}
