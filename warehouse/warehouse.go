package warehouse

import "fmt"

type Warehouse struct {
	name         string
	products     map[int]*product
	suppliers    map[int]*supplier
	supplyOrders map[int]*order
}

func NewWarehouse(name string) *Warehouse {
	if name == "" {
		fmt.Println("")
		fmt.Println("Ошибка : Имя должно быть заполнено")
		return nil
	}

	return &Warehouse{name: name, products: make(map[int]*product), suppliers: make(map[int]*supplier), supplyOrders: make(map[int]*order)}
}

func (w *Warehouse) AddProduct(id int, name string, stock int, price float64) {
	if _, ok := w.products[id]; ok {
		fmt.Println("")
		fmt.Println("Ошибка : Продукт с таким ID уже есть на складе")
		return
	}

	w.products[id] = newProduct(id, name, stock, price)
}

func (w *Warehouse) GetProducts() {
	if len(w.products) <= 0 {
		fmt.Println("")
		fmt.Println("Список продуктов на складе пуст")
		return
	}

	fmt.Println("")
	fmt.Println("Список продуктов на складе:")

	for _, p := range w.products {
		fmt.Println("ID:", p.id, "; Название:", p.name, "; Количество:", p.stock, "; Цена:", p.price)
	}

}

func (w *Warehouse) AddSupllier(id int, name string) {
	if _, ok := w.suppliers[id]; ok {
		fmt.Println("")
		fmt.Println("Ошибка : Поставщик с таким ID уже существует в системе")
		return
	}

	w.suppliers[id] = newSupplier(id, name)
}

func (w *Warehouse) GetSuppliers() {
	if len(w.suppliers) <= 0 {
		fmt.Println("")
		fmt.Println("Список поставщиков пуст")
		return
	}

	fmt.Println("")
	fmt.Println("Список поставщиков:")

	for _, s := range w.suppliers {
		fmt.Println("ID:", s.id, "; Поставщик:", s.name)
	}
}

func (w *Warehouse) CreateSupplyOrder(orderID, supplierID int) {
	if orderID <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : ID заказа должен быть больше нуля")
		return
	}

	if supplierID <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : ID поставщика должен быть больше нуля")
		return
	}

	if _, ok := w.suppliers[supplierID]; !ok {
		fmt.Println("")
		fmt.Println("Ошибка : Поставщик с таким ID не найден")
		return
	}

	if _, ok := w.supplyOrders[orderID]; ok {
		fmt.Println("")
		fmt.Println("Ошибка : Поставка с таким ID уже создана")
		return
	}

	w.supplyOrders[orderID] = newOrder(orderID, w.suppliers[supplierID])
}

func (w *Warehouse) GetSupplyOrders() {
	if len(w.supplyOrders) <= 0 {
		fmt.Println("")
		fmt.Println("Список поставок пуст")
		return
	}

	fmt.Println("")
	fmt.Println("Список поставок:")

	for _, o := range w.supplyOrders {
		fmt.Println("ID:", o.id, "; Статус:", o.status)
		fmt.Println("Продукты:", o.products, "; Поставщик:", o.supplier.name)
		fmt.Println("-------------")
	}
}

func (w *Warehouse) AddProductToOrder(orderID, productID, qty int) {
	if orderID <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : ID заказа должен быть больше нуля")
		return
	}

	if productID <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : ID продукта должен быть больше нуля")
		return
	}

	if qty <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : Количество товаров должно быть больше нуля")
		return
	}

	if _, ok := w.supplyOrders[orderID]; !ok {
		fmt.Println("")
		fmt.Println("Ошибка : Поставка с таким ID не найдена")
		return
	}

	if w.supplyOrders[orderID].status == ACCEPTED {
		fmt.Println("")
		fmt.Println("Ошибка : Нельзя добавить товар в подтвержденную поставку")
		return
	}

	if _, ok := w.products[productID]; !ok {
		fmt.Println("")
		fmt.Println("Ошибка : Продукт с таким ID не найден")
		return
	}

	productsOfOrder := w.supplyOrders[orderID].products

	productsOfOrder[w.products[productID]] = qty

}

func (w *Warehouse) ConfirmSupplyOrder(orderID int) {
	if orderID <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : ID заказа должен быть больше нуля")
		return
	}

	if _, ok := w.supplyOrders[orderID]; !ok {
		fmt.Println("")
		fmt.Println("Ошибка : Поставка с таким ID не найдена")
		return
	}

	if len(w.supplyOrders[orderID].products) <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : Нельзя подтвердить пустую поставку")
		return
	}

	if w.supplyOrders[orderID].status == ACCEPTED {
		fmt.Println("")
		fmt.Println("Ошибка : Заказ уже подтвержден")
		return
	}

	if w.supplyOrders[orderID].status == CANCELED {
		fmt.Println("")
		fmt.Println("Ошибка : Заказ отменен")
		return
	}

	productsOfOrder := w.supplyOrders[orderID].products

	for p, q := range productsOfOrder {
		p.stock += q
	}
	w.supplyOrders[orderID].status = ACCEPTED
}

func (w *Warehouse) CancelSupplyOrder(orderID int) {
	if orderID <= 0 {
		fmt.Println("")
		fmt.Println("Ошибка : ID заказа должен быть больше нуля")
		return
	}

	if _, ok := w.supplyOrders[orderID]; !ok {
		fmt.Println("")
		fmt.Println("Ошибка : Поставка с таким ID не найдена")
		return
	}

	if w.supplyOrders[orderID].status == CANCELED {
		fmt.Println("")
		fmt.Println("Ошибка : Нельзя отменить и так отмененный заказ")
		return
	}

	if w.supplyOrders[orderID].status == ACCEPTED {
		fmt.Println("")
		fmt.Println("Ошибка : Нельзя отменить подтвержденный заказ")
		return
	}

	w.supplyOrders[orderID].status = CANCELED
}
