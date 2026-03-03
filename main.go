package main

import (
	"warehouse-system/warehouse"
)

func main() {

	centralWarehouse := warehouse.NewWarehouse("Центральный Склад")

	centralWarehouse.AddProduct(1, "Яблоко", 3, 40.0)
	centralWarehouse.AddProduct(2, "Бананы", 4, 67.0)
	centralWarehouse.AddProduct(3, "Лайм", 5, 55.0)

	centralWarehouse.AddSupllier(1, "ООО Бнал")
	centralWarehouse.AddSupllier(2, "ООО Бнал")

	centralWarehouse.CreateSupplyOrder(1, 1)

	centralWarehouse.GetSupplyOrders()

	centralWarehouse.AddProductToOrder(1, 2, 5)

	centralWarehouse.GetSupplyOrders()

	centralWarehouse.ConfirmSupplyOrder(1)

	centralWarehouse.GetSupplyOrders()

	centralWarehouse.CancelSupplyOrder(1)

}
