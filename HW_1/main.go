package main

import (
	"fmt"
)

type Order struct {
	ID       int
	Customer *OrderCustomer
	Goods    []*OrderGood
}

type OrderCustomer struct {
	name     string
	surname  string
	phoneNum string
}

type OrderGood struct {
	name  string
	price float64
	count int
}

var Orders = make(map[int]*Order)

func main() {

	// Создать покупателя
	customer1 := OrderCustomer{"Tom", "Black", "8-984-156-19-66"}

	// Создать заказ на покупателя
	order := NewOrder(&customer1)

	// Добавить в заказ 3 произвольных товара
	AddGood(order, "opium", 10.0, 2)
	AddGood(order, "cola", 2.0, 3)
	AddGood(order, "snaks", 1.0, 4)

	// Вывести заказ на экран
	fmt.Println("ToString")
	fmt.Println(ToString(order))

	// Повторить с п.1 - создать еще один заказ
	customer2 := OrderCustomer{"Adam", "Drown", "8-999-156-19-66"}
	_ = NewOrder(&customer2)

	// Поптаться получить заказ по ID и вывести его на экран
	fmt.Println("GetOrder(1)")
	o1, err1 := GetOrderByID(1)
	if err1 == nil {
		fmt.Println(ToString(o1))
	} else {
		fmt.Println("Incorrect Order ID")
	}
	o2, err2 := GetOrderByID(2)
	if err2 == nil {
		fmt.Println(ToString(o2))
	} else {
		fmt.Println("Incorrect Order ID")
	}
	// Попытаться получить несуществующий заказ и вывести его на экран
	o3, err3 := GetOrderByID(3)
	if err3 == nil {
		fmt.Println(ToString(o3))
	} else {
		fmt.Println("Incorrect Order ID")
	}
	// for k, v := range Orders {
	// 	fmt.Println(k, ToString(v))
	// }
}

func NewOrder(customer *OrderCustomer) *Order {
	var NewOrder = Order{
		ID:       len(Orders) + 1,
		Customer: customer,
	}

	Orders[NewOrder.ID] = &NewOrder

	return &NewOrder
}

func AddGood(order *Order, name string, price float64, count int) {
	order.Goods = append(order.Goods, &OrderGood{name, price, count})

	Orders[order.ID] = order
}

func ToString(order *Order) string {
	var OrderGoodsToString string
	for i := range order.Goods {
		OrderGoodsToString += fmt.Sprintf("\t%d name: %s price: %f count: %d \n", i+1, order.Goods[i].name, order.Goods[i].price, order.Goods[i].count)
	}

	OrderToString := fmt.Sprintf("ID: %d\nName: %s\nSurname: %s\nPhone: %s\nGoods:\n%s", order.ID, order.Customer.name, order.Customer.surname, order.Customer.phoneNum, OrderGoodsToString)

	return OrderToString
}

func GetOrderByID(ID int) (*Order, error) {
	if v, ok := Orders[ID]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("incorrect order id")
}
