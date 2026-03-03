package warehouse

import "fmt"

type product struct {
	id    int
	name  string
	stock int
	price float64
}

func newProduct(id int, name string, stock int, price float64) *product {
	if id <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : ID должен быть больше нуля")
		return nil
	}

	if name == "" {
		fmt.Println("")
		fmt.Println("Ошибка : Название товара должно быть заполнено")
		return nil
	}

	if stock < 0 {
		fmt.Println("")
		fmt.Println("Ошибка : Количество товаров не может быть отрицательное")
		return nil
	}

	if price <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : Цена должна быть выше нуля")
		return nil
	}

	return &product{id: id, name: name, stock: stock, price: price}
}
