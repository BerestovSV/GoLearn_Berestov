package orderRepository

import (
	"fmt"
	"main/orderService/order"
)

var orders = make(map[int]order.Order)

///////////////////////////////////// Public always in top

// создание НОВОГО заказа
func Store(order order.Order) error {

	order.ID = findMaxKey() + 1
	orders[order.ID] = order

	return nil
}

// удаление записей из мапы
func Delete(id int) error {
	return nil
}

// поиск по ID
func Get(id int) (*order.Order, error) {
	if v, ok := orders[id]; ok {
		return &v, nil
	}
	return nil, fmt.Errorf("incorrect order id")
}

// обновление существующего заказа
func Update(order order.Order) error {
	return nil
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
