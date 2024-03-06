package order

import (
	"fmt"
	"main/orderService/order/orderCustomer"
	"main/orderService/order/orderGood"
)

type Order struct {
	ID       int
	Customer orderCustomer.OrderCustomer
	Goods    []orderGood.OrderGood
	Status   string
}

func New(customer orderCustomer.OrderCustomer) Order {
	return Order{
		Customer: customer,
		Status:   initialized,
	}
}

func (order *Order) AddGood(name string, price float64, count int) {
	if !isEditable(order.Status) {
		fmt.Println("\nЗаказ выполнен, редактирование невозможно")
	} else {
		order.Goods = append(order.Goods, orderGood.New(name, price, count))
	}
}

func (order *Order) String() string {
	var orderGoodsToString string
	var orderSum float64 = 0
	for i := range order.Goods {
		orderGoodsToString += fmt.Sprintf("\t%d name: %s price: %f count: %d \n", i+1, order.Goods[i].Name, order.Goods[i].Price, order.Goods[i].Count)
		orderSum += order.Goods[i].Price * float64(order.Goods[i].Count)
	}

	orderToString := fmt.Sprintf("ID: %d\nName: %s\nSurname: %s\nPhone: %s\nGoods:\n%s \nTotal price: %f \nStatus: %s",
		order.ID, order.Customer.Name, order.Customer.Surname, order.Customer.PhoneNum, orderGoodsToString, orderSum, order.Status)

	return orderToString
}

// удаление части товаров из заказа
func (order *Order) DeleteGood(flag byte, goodNum int) {
	if !isEditable(order.Status) {
		fmt.Println("\nЗаказ выполнен, редактирование невозможно")
	} else {
		switch flag {
		case 0:
			// сначала проверяем, что товарная позиция не выходит за пределы слайса, если все "ок", то выполняем удаление записи
			if goodNum <= len(order.Goods) && goodNum >= 0 {
				order.Goods = append(order.Goods[0:goodNum-1], order.Goods[goodNum:]...)
			}

		case 1:
			emptyGoodSlice := []orderGood.OrderGood{}
			order.Goods = emptyGoodSlice
		}
	}
}

func (order *Order) Init() {
	order.Status = completed
}

// 	// обработка ошибок, для справки
// 	fmt.Println("GetOrder(1)")
// 	o1, err1 := GetOrderByID(1)
// 	if err1 == nil {
// 		fmt.Println(ToString(o1))
// 	} else {
// 		fmt.Println("Incorrect Order ID")
// 	}
