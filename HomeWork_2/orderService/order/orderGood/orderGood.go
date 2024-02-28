package orderGood

type OrderGood struct {
	name  string
	price float64
	count int
}

func New(name string, price float64, count int) OrderGood {
	return OrderGood{
		name:  name,
		price: price,
		count: count,
	}
}
