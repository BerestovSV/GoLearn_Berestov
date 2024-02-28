package order

import (
	"main/orderService/order/orderCustomer"
	"main/orderService/order/orderGood"
)

// вывести каждую структуру в свой пакет
type Order struct {
	ID       int
	Customer orderCustomer.OrderCustomer
	Goods    []orderGood.OrderGood
}

func New(customer orderCustomer.OrderCustomer) Order {
	return Order{
		Customer: customer,
	}
}

func (order *Order) AddGood(name string, price float64, count int) {
	order.Goods = append(order.Goods, orderGood.New(name, price, count))
}
