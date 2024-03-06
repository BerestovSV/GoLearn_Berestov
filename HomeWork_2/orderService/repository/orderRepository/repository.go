package orderRepository

import (
	"fmt"
	"main/orderService/order"
)

var orders = make(map[int]order.Order)

///////////////////////////////////// Public always in top

// создание НОВОГО заказа
func Store(order *order.Order) error {
	// переделать под обновление заказа
	order.ID = findMaxKey() + 1
	order.Init()
	orders[order.ID] = *order

	return nil
}

// удаление заказа
func Delete(id int) error {
	delete(orders, id)
	return fmt.Errorf("incorrect order id")
}

// поиск по ID
func Get(id int) (*order.Order, error) {
	if v, ok := orders[id]; ok {
		return &v, nil
	}
	return nil, fmt.Errorf("incorrect order id")
}

///////////////////////////////////// private always under Public

// находим максимальный ключ
func findMaxKey() int {
	maxKey := 0
	for k := range orders {
		if k > maxKey {
			maxKey = k
		}
	}
	return maxKey
}
