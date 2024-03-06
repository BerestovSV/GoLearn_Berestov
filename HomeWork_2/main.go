package main

import (
	"fmt"
	"main/orderService/order"
	"main/orderService/order/orderCustomer"
	"main/orderService/repository/orderRepository"
)

func main() {

	// 1.	Создать заказ с покупателем и добавить в него товары
	customer1 := orderCustomer.New("Tom", "Black", "8-984-156-19-66")

	order1 := order.New(customer1)

	order1.AddGood("opium", 10.0, 2)
	order1.AddGood("cola", 2.0, 3)
	order1.AddGood("snaks", 1.0, 4)
	fmt.Println("1\n", order1.String())

	// 2.	Удалить часть товаров: Если первый аргумент "1", то удаляем все товары, если "0", то удаляется товар по указанной позиции
	order1.DeleteGood(0, 3)
	fmt.Println("\n2\n", order1.String())

	// 3.	Выполнить заказ
	orderRepository.Store(&order1)
	fmt.Println("\n3\n", order1.String())

	// 4.	Попытаться добавить товар в заказ - обработать ошибку и вывести на экран
	order1.AddGood("spice", 4.0, 10)

	// 5.	Попытаться удалить товар из заказа - также обработать ошибку
	order1.DeleteGood(1, 0)
}
